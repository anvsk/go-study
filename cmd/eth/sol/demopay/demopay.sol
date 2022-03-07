// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.4;

//   - n 个人消费(比如吃饭)之后使用合约进行支付账单
//   - 所有人都可以查看账单
//   - 任何一个人都可以发起账单
//   - 一个人确认账单即可进行支付
//   - 合约中支付账单的资金不用考虑谁来出

contract DemoPay {
    // 老板账户\用于收款
    address boss;

    ////// 购物车的商品操作  [1.construct 创建商品]

    // 商品
    struct Goods {
        uint256 id;
        string name;
        uint256 price;
    }

    // 商品集合
    mapping(uint256 => Goods) allgoods;

    // 单项明细[1. 添加商品]
    struct GoodsRow {
        Goods goods;
        uint256 nums;
    }

    // 购物车(类似房间号、桌号)[1.领头人开车 2.其他人可上车]
    struct Car {
        uint256 id;
        GoodsRow[] goodsdetail; // 买的东西
        address[] members; //上车的人
    }

    // carid=>car购物车集合
    mapping(uint256 => Car) allcars;
    // uaddr=>car用户集合
    mapping(address => Car) allusers;

    ////// 发起账单、确认账单

    // 账单[1. 发起账单(待确认) 2. 确认账单(待付款) 3. 付款]
    struct Order {
        uint256 carnum;
        Car cardetail;
        uint256 amount;
        bool confirmed;
        bool payed;
    }

    // 订单集合
    mapping(uint256 => Order) allorders;

    // 其他方法 [1. 查看账单]

    /**** function *****/

    constructor(Goods[] memory goods, address bossaddress) {
        for (uint256 index = 0; index < goods.length; index++) {
            allgoods[goods[index].id] = goods[index];
        }
        boss = bossaddress;
    }

    // 开车
    function newCar() external returns (uint256 carnum) {
        // todo 生成随机数
        carnum = 1;
        allcars[carnum].id = carnum;
        allcars[carnum].members.push(msg.sender);
        allusers[msg.sender] = allcars[carnum];
    }

    // 上车
    function takeCar(uint256 carnum) external {
        require(allcars[carnum].id > 0, "car nums error");
        allcars[carnum].members.push(msg.sender);
        allusers[msg.sender] = allcars[carnum];
    }

    // 添加商品
    function addGoods(uint256 gid, uint256 nums) external {
        uint256 carnum = getCarNum();
        require(allgoods[gid].id > 0, "goods id error");
        require(nums > 0, " nums must gt zero ");
        allcars[carnum].members.push(msg.sender);
        allcars[carnum].goodsdetail.push(
            GoodsRow({goods: allgoods[gid], nums: nums})
        );
    }

    // 发起账单
    function getOrder() public {
        uint256 carnum = getCarNum();
        require(allorders[carnum].carnum == 0, "order had getted");
        allorders[carnum].carnum = carnum;
        allorders[carnum].cardetail = allcars[carnum];
        uint256 amount = 0;
        for (
            uint256 index = 0;
            index < allcars[carnum].goodsdetail.length;
            index++
        ) {
            // todo
            // amount +=
            //     allcars[carnum].goodsdetail[index].nums *
            //     amount += allcars[carnum].goodsdetail[index].goods.price;
            amount += allcars[carnum].goodsdetail[index].goods.price;
        }
        allorders[carnum].amount = amount;
    }

    // 确认账单
    function confirmOrder() external {
        uint256 carnum = getCarNum();
        require(allorders[carnum].carnum > 0, "order has not getted");
        allorders[carnum].confirmed = true;
    }

    // 付款账单
    function payOrder() external {
        uint256 carnum = getCarNum();
        require(allorders[carnum].carnum > 0, "order is not exsist");
        allorders[carnum].payed = true;
        payable(boss).transfer(allorders[carnum].amount);
    }

    // 账单详情
    function showOrder() public view returns (Order memory order) {
        uint256 carnum = getCarNum();
        require(allorders[carnum].carnum > 0, "order isnot exsit");
        order = allorders[carnum];
    }

    // 获取房间号
    function getCarNum() public view returns (uint256 carnum) {
        Car memory car = allusers[msg.sender];
        require(car.id > 0, "carnum is error");
        carnum = car.id;
    }
}
