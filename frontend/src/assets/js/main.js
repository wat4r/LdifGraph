
var groups = ['null'];
var users = ['null'];
var type_value = 'group';
var dom = document.getElementById('container');
var ldfiChart = echarts.init(dom, 'dark', {
    renderer: 'canvas',
    useDirtyRect: true
});
var app = {};
var graph_data = {
    "name": "Ldif Graph",
    "children": null
}

var value_select_mdui = new mdui.Select('#value-select');
var value_select = document.getElementById('value-select');
var type_select = document.getElementById('type-select');

function set_options(update) {
    let select_type = type_select.value;
    let items;

    if (select_type == type_value && !update) {
        return;
    }

    switch (select_type) {
        case 'group':
            items = groups;
            break;
        case 'user':
            items = users;
            break;
        default:
            return;
    }

    type_value = select_type;

    value_select.innerHTML = '';
    items.forEach(item => {
        var opt = document.createElement("option");
        opt.value = item;
        opt.text = item;
        value_select.add(opt);
    });

    value_select_mdui.handleUpdate();
}

set_options(true);

window.runtime.EventsOn("set_groups", function (data) {
    groups = data;
    set_options(true);
})

window.runtime.EventsOn("set_users", function (data) {
    users = data;
    set_options(true);
})

mdui.$('#type-select').on('closed.mdui.select', function (e) {
    set_options(false);
});

window.runtime.EventsOn("send_msg", function (msg) {
    mdui.snackbar({
        message: msg,
        position: 'right-top',
        buttonColor: 'red'
    });
})

var rendering_btn = document.getElementById('rendering');
rendering_btn.addEventListener("click", (event) => {
    let select_type = type_select.value;
    let select_value = value_select.value;
    if (select_value != "null") {
        ldfiChart.showLoading();
        let res = window.go.main.App.GetGraphData(select_type, select_value);
        res.then((value) => {
            let data = JSON.parse(value);
            refreshData(data);
            ldfiChart.hideLoading();
        }).catch((error) => {
            ldfiChart.hideLoading();
            mdui.snackbar({
                message: error,
                position: 'right-top',
                buttonColor: 'red'
            });
        });
    }
});

ldfiChart.showLoading();
var ldfiOption = {
    tooltip: {
        trigger: 'item',
        triggerOn: 'mousemove'
    },
    legend: {
        data: [
            {
                name: 'Group',
                icon: 'emptyCircle'
            },
            {
                name: 'User',
                icon: 'roundRect'
            }
        ],
        icon: 'rect',
        textStyle: {
            color: '#ccc'
        }
    },
    series: [
        {
            type: 'tree',
            data: [graph_data],
            top: 'middle',
            left: 'center',
            layout: 'orthogonal',
            symbol: function (value, params) {
                let svg = "roundRect";
                if (value == 1) {
                    svg = "emptyCircle";
                }
                return svg
            },
            symbolSize: 12,
            roam: true,
            label: {
                color: 'inherit',
                fontFamily: 'sans-serif',
                fontSize: 14,
                position: 'top',
                verticalAlign: 'middle'
            },
            leaves: {
                label: {
                    position: 'right',
                    verticalAlign: 'middle',
                    align: 'left'
                }
            },
            emphasis: {
                focus: 'relative'
            },
            expandAndCollapse: true,
            animationDuration: 550,
            animationDurationUpdate: 750
        }
    ]
}


ldfiOption && ldfiChart.setOption(ldfiOption, true);
ldfiChart.hideLoading();

window.addEventListener('resize', ldfiChart.resize);

function refreshData(data) {
    console.log(data);
    ldfiOption.series[0].data = data;
    ldfiChart.setOption(ldfiOption, true);
}