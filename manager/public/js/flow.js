var FlowClass  = {
    ProjectId: 0,
    AgetLength: "tenminute",
    CanvasId: "",
    ChartType: "line",
    DisplayFormat: "increment",//increment,full;

    CountSum:0,
    ByteSizeSum:0,
    CallBack:null,
    Data:[],

    maxCount:0,
    maxByteSize:0,

    setProjectId: function (id) {
        this.ProjectId = id;
    },
    setAgetLength: function (AgetLength) {
        this.AgetLength = AgetLength;
    },
    setCanvasId: function (CanvasId) {
        this.CanvasId = CanvasId;
    },

    setDisplayFormat: function (DisplayFormat) {
        this.DisplayFormat = DisplayFormat;
    },

    setCallBackFun: function (f) {
        if (typeof(f) == "function"){
            this.CallBack = f;
        }
    },

    getData:function () {
        return this.Data;
    },

    add0: function (m) {
        return m < 10 ? '0' + m : m
    },

    TimeFormat: function (timeUnix) {
        var time = new Date(parseInt(timeUnix) * 1000);
        var y = time.getFullYear();
        var m = time.getMonth();
        var d = time.getDate();
        var h = time.getHours();
        var mm = time.getMinutes();
        var s = time.getSeconds();
        return this.add0(h) + ':' + this.add0(mm) + ':' + this.add0(s);
        //return y + '-' + this.add0(m) + '-' + this.add0(d) + ' ' + this.add0(h) + ':' + this.add0(mm) + ':' + this.add0(s);
    },

    init_data: function () {
        return {
            color: ["#1ab394", "#5CACEE"],
            tooltip: {
                trigger: "axis"
            },
            legend: {
                data: ["subscribers_count", "stargazers_count", "forks_count"]
            },
            calculable: !0,
            xAxis: [{
                type: "category",
                boundaryGap: !1,
                data: []
            }],
            yAxis: [{
                type: "value"
            }],
            series: [
                {
                    name: "subscribers_count",
                    type: "line",
                    data: [],
                    markPoint: {
                        data: [{
                            type: "max",
                            name: "最大值"
                        },
                            {
                                type: "min",
                                name: "最小值"
                            }]
                    }

                },
                {
                    name: "stargazers_count",
                    type: "line",
                    data: [],
                    markPoint: {
                        data: [{
                            type: "max",
                            name: "最大值"
                        },
                            {
                                type: "min",
                                name: "最小值"
                            }]
                    }
                },
                {
                    name: "forks_count",
                    type: "line",
                    data: [],
                    markPoint: {
                        data: [{
                            type: "max",
                            name: "最大值"
                        },
                            {
                                type: "min",
                                name: "最小值"
                            }]
                    }
                }]
        };
    },

    rewrite_data: function (d) {
        if ($("#" + this.CanvasId).length <= 0) {
            return
        }
        if (d.length == 0) {
            this.Data = [];
            return false
        }
        this.Data = d;

        var e = echarts.init(document.getElementById(this.CanvasId));
        var a = this.init_data();

        for( var i in d){
            a.xAxis[0].data.push(d[i].Add_time);
            a.series[0].data.push(d[i].Subscribers_count);
            a.series[1].data.push(d[i].Stargazers_count);
            a.series[2].data.push(d[i].Forks_count);
        }
        e.setOption(a);
        $(window).resize(e.resize);
        d = null;

    },

    incrementData: function (d) {
        var data = [];
        var lasttime = -1;
        var Subscribers_count = 0
        var Stargazers_count = 0
        var Forks_count = 0
        for (s in d) {
            if (lasttime > 0) {
                var tSubscribers_count = d[s].Subscribers_count - Subscribers_count;
                if (tSubscribers_count < 0) {
                    tSubscribers_count = 0;
                }
                var tStargazers_count = d[s].Stargazers_count - Stargazers_count;
                if (tStargazers_count < 0) {
                    tStargazers_count = 0;
                }
                var tForks_count = d[s].Forks_count - Forks_count;
                if (tForks_count < 0) {
                    tForks_count = 0;
                }
                data.push({
                    time: this.TimeFormat(d[s].Add_time),
                    Subscribers_count: tSubscribers_count,
                    Stargazers_count: tStargazers_count,
                    Forks_count: tForks_count,
                });
                Subscribers_count = d[s].Subscribers_count;
                Stargazers_count = d[s].Stargazers_count;
                Forks_count = d[s].Forks_count;
            }else{
                Subscribers_count = 0;
                Stargazers_count = 0;
                Forks_count = 0;
            }
            lasttime = d[s].Add_time;
        }
        return data;
    },

    fullData: function (d) {
        var data = [];
        for (s in d) {
            if (d[s].Add_time != "") {
                data.push({
                    time: this.TimeFormat(d[s].Add_time),
                    Subscribers_count: d[s].Subscribers_count,
                    Stargazers_count: d[s].Stargazers_count,
                    Forks_count: d[s].Forks_count,
                });
            }
        }
        return data;
    },

    getFlowData: function () {
        var obj = this;
        this.ByteSizeSum = 0;
        this.CountSum = 0;
        if ( obj.ProjectId <= 0 ){
            return;
        }
        $.post(
            "/flow/get",
            {
                project_id: obj.ProjectId,
            },
            function (d, status) {
                if (status != "success") {
                    return false;
                }

                if (obj.DisplayFormat == "full") {
                    obj.rewrite_data(obj.fullData(d));
                } else {
                    obj.rewrite_data(obj.incrementData(d));
                }
                if(obj.CallBack != null){
                    obj.CallBack();
                }
            }, 'json');
    },
}