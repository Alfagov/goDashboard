htmx.defineExtension('echarts', {
    transformResponse: function (text, xhr, elt) {
        // parse json data
        var data = JSON.parse(text);

        // fetch echart element
        var option = data;
        chartsMap.get(data.id).chart.setOption(option);

        return "";
    }
});