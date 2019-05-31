# WS-USB01-THP data read library for using golang

The "thp" is WS-USB01-THP data read library with using PlanexCloud. You can get clearly JSON data with this.

## how to build example

`$ go build example/thp.go`

## how to use example

You have to get device's MAC address and TOKEN on [Planex web site](http://svcipp.planex.co.jp/iot/wp-login.php) at first.

`$ DOCODEMO_MAC="xx:yy:zz:mm:nn:oo" DOCODEMO_TOKEN="xxxxxxx" ./thp`
    
example output:
>[{"date":"2019-05-3102:18:02","thermal":28.18,"humidity":51.63,"pressure":1011.6},{"date":"2019-05-3102:18:25","thermal":28.18,"humidity":51.54,"pressure":1011.83},{"date":"2019-05-3102:18:48","thermal":28.18,"humidity":51.57,"pressure":1011.36},{"date":"2019-05-3102:19:10","thermal":28.18,"humidity":51.64,"pressure":1011.6},{"date":"2019-05-3102:19:34","thermal":28.18,"humidity":51.67,"pressure":1011.6},{"date":"2019-05-3102:19:56","thermal":28.18,"humidity":51.59,"pressure":1011.6},{"date":"2019-05-3102:20:20","thermal":28.18,"humidity":51.58,"pressure":1011.13},{"date":"2019-05-3102:20:43","thermal":28.18,"humidity":51.55,"pressure":1011.36},{"date":"2019-05-3102:21:06","thermal":28.15,"humidity":51.65,"pressure":1011.52},{"date":"2019-05-3102:21:28","thermal":28.15,"humidity":51.61,"pressure":1011.29},{"date":"2019-05-3102:21:51","thermal":28.15,"humidity":51.6,"pressure":1011.75},{"date":"2019-05-3102:22:14","thermal":28.15,"humidity":51.7,"pressure":1011.75},{"date":"2019-05-3102:22:37","thermal":28.13,"humidity":51.68,"pressure":0}]

## known issue

The timestamp on golang Struct from JSON is broken. It missed parsing timestamp data, I'm sure I'll fix it soon but currently, I don't need that data, so I may leave it for a while...

