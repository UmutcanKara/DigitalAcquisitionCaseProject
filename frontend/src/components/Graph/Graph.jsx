import React from "react";
import Chart from "react-apexcharts";
import {Box, CircularProgress} from "@mui/material";

const Graph = ({xaxis ,yaxis ,isLoading}) => {
    const [chartData, setChartData] = React.useState({
        options: {
            chart: {
                id: "line-graph"
            },
            xaxis: {
                categories: []
            }
        },
        series: [
            {
                name: "temps",
                data: []
            }
        ]
    })
    React.useEffect(() => {
        if (!xaxis || !yaxis) return
        let length = xaxis.length
        let step = parseInt((length / 100).toFixed(0))
        setChartData(prevState => ({
            options: {
                ...prevState.options,
                xaxis: {
                    categories: xaxis.filter((_, idx) => idx % step === 0).map(isoTime => {
                        let parsedDate = new Date(isoTime)
                        let splitDate = parsedDate.toString().split(" ")
                        // Return the month & year
                        return `${splitDate[1]} ${splitDate[2]} ${splitDate[3]}`
                    })
                }
            },
            series: [
                {
                    ...prevState.series,
                    data: yaxis.filter((_, idx) => idx % step === 0)
                }
            ]
        }))
    }, [xaxis, isLoading])


    return <Box display="flex" width="100%" justifyContent="center" alignItems="center" flexDirection="column">
            {!isLoading ? <Chart
                options={chartData.options}
                series={chartData.series}
                type="line"
                width="500"
            /> : <Box><CircularProgress/></Box>}
        </Box>

}

export default Graph