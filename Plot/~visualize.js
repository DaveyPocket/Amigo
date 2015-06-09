$(function() {
    var dataFileNum = window.location.search.replace("?", "");
    $.get("data"+dataFileNum+".json", function(file) {
        $("#graph").highcharts({
            chart: {
                zoomType: 'x'
            },
            legend: {
                enabled: false
            },
            credits: {
                enabled: false
            },
            title: {
                text: ''
            },
            series: [{
              data: file.fft,
              yAxis: 0
            }],
          yAxis: [{ title: { text: "FFT" } },
                  { title: { text: "Signal" }, opposite: true }]
        });
        $("#text").html(file.text);
    });
});
