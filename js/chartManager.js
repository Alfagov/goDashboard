var chartsMap = new Map();

function addChart(chart, chartId, element) {

    var observer = new ResizeObserver(() => chart.resize());
    observer.observe(element);

    chartsMap.set(chartId, {
        chart: chart,
        observer: observer
    });
}

function disposeAllCharts() {
    for (let item of chartsMap.values()) {
        item.chart.dispose();
        item.observer.disconnect();
    }
    chartsMap.clear();
}