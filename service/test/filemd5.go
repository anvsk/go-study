package test

import (
    "crypto/md5"
    "errors"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "sort"
    "sync"
    "time"

    "github.com/pieterclaerhout/go-log"
)

func Bootstrap() {

    m, err := MD5All2(os.Args[1])
    log.Debug("xxxx")
    if err != nil {
        log.Debug("error")
        fmt.Println(err)
        panic(err)
    }
    log.Debug("noerror")

    var paths []string
    for path := range m {
        paths = append(paths, path)
    }
    sort.Strings(paths)
    for _, path := range paths {
        log.Debugf("%x  %s\n", m[path], path)
    }
}

type result struct {
    path string
    sum  [md5.Size]byte
    err  error
}

func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
    // For each regular file, start a goroutine that sums the file and sends
    // the result on c.  Send the result of the walk on errc.
    c := make(chan result)
    errc := make(chan error, 1)
    go func() {
        var wg sync.WaitGroup
        err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if !info.Mode().IsRegular() {
                return nil
            }
            wg.Add(1)
            go func() {
                data, err := ioutil.ReadFile(path)
                // select {
                // case c <- result{path, md5.Sum(data), err}:
                // case <-done:
                // }
                c <- result{path, md5.Sum(data), err}
                wg.Done()
            }()
            // Abort the walk if done is closed.
            select {
            case <-done:
                return errors.New("walk canceled")
            default:
                return nil
            }
        })
        // Walk has returned, so all calls to wg.Add are done.  Start a
        // goroutine to close c once all the sends are done.
        go func() {

            wg.Wait()
            log.Debug("bb")

            close(c)
        }()
        log.Debug("cc")

        // No select needed here, since errc is buffered.
        errc <- err
    }()
    return c, errc
}

func MD5All(root string) (map[string][md5.Size]byte, error) {
    // MD5All closes the done channel when it returns; it may do so before
    // receiving all the values from c and errc.
    done := make(chan struct{})
    defer close(done)

    c, errc := sumFiles(done, root)

    m := make(map[string][md5.Size]byte)
    for r := range c {
        if r.err != nil {
            return nil, r.err
        }
        m[r.path] = r.sum
    }
    if err := <-errc; err != nil {
        return nil, err
    }
    return m, nil
}

func MD5All2(root string) (map[string][md5.Size]byte, error) {
    log.Debug("do2222file")

    // MD5All closes the done channel when it returns; it may do so before
    // receiving all the values from c and errc.
    done := make(chan struct{})
    defer close(done)
    m := make(map[string][md5.Size]byte)

    paths, errc := walkFiles(done, root)

    r := dojobfile(paths, done)

    log.Debug("333332Afterdofile")

    for v := range r {
        if v.err != nil {
            return m, v.err
        }
        m[v.path] = v.sum
    }

    if err := <-errc; err != nil {
        return nil, err
    }

    return m, nil
}

func dojobfile(pathsCh <-chan string, done chan struct{}) chan result {
    res := make(chan result)
    tasks := 10
    wg := sync.WaitGroup{}
    handle := func(pathsCh <-chan string) {
        defer wg.Done()
        for {
            select {
            case path, ok := <-pathsCh:
                if !ok {
                    return
                } else {
                    time.Sleep(30 * time.Millisecond)
                    data, err := ioutil.ReadFile(path)
                    res <- result{path: path, sum: md5.Sum(data), err: err}
                }
            case <-done:
                log.Info("comming dojob!")
                return
            }
        }
    }
    wg.Add(tasks)
    for i := 0; i < tasks; i++ {
        go handle(pathsCh)
    }
    go func() {
        wg.Wait()
        close(res)
        // close(done)
    }()

    return res
}

func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
    paths := make(chan string)
    errc := make(chan error, 1)
    go func() {
        // Close the paths channel after Walk returns.
        defer close(paths)
        // No select needed for this send, since errc is buffered.
        errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if !info.Mode().IsRegular() {
                return nil
            }
            select {
            case paths <- path:
            case <-done:
                return errors.New("walk canceled")
            }
            return nil
        })
    }()
    return paths, errc
}

// func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
//     for path := range paths {
//         data, err := ioutil.ReadFile(path)
//         select {
//         case c <- result{path, md5.Sum(data), err}:
//         case <-done:
//             return
//         }
//     }
// }
