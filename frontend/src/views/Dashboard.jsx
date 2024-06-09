import React from "react";
import {WeatherService} from "../api/service/weather.js";
import {AuthContext} from "../context/AuthContext.jsx";
import {AuthService} from "../api/service/auth.js";
import {useNavigate} from "react-router-dom";
import Chart from "react-apexcharts";
import {Box, Button, ButtonGroup, CircularProgress, Typography} from "@mui/material";

const Dashboard = () => {
    const [isLoading, setIsLoading] = React.useState(true)
    const [userData, setUserData] = React.useState({
        username: "",
        hometown: ""
    })
    const [chartData, setChartData] = React.useState()
    const navigate = useNavigate()
    const { auth, logout } = React.useContext(AuthContext)
    React.useEffect(() => {
        const getUserData = async () => {
            let user = await AuthService.getUser(auth.username)
            if (user === "Unauthorized") {
                logout()
                navigate("/")
                return
            }
            user = user.data
            setUserData(user)
            await fetchWeatherData(user)

        }
        getUserData().then().catch(err => console.error(err))

    }, [])

    const fetchWeatherData = async (user, depthChoice = "3M") => {
        setIsLoading(true)
        let now = new Date()
        switch (depthChoice) {
            case "3M":
                now.setMonth(now.getMonth() -3)
                break;
            case "1Y":
                now.setFullYear(now.getFullYear() -1)
                break;
            case "5Y":
                now.setFullYear(now.getFullYear() -5)
                break;

        }
        let isoDate = now.toISOString().split("T")[0]
        let weatherData = await WeatherService.getWeather(user.hometown ,isoDate)
        let length = weatherData.data.hourly.time.length
        // 100 = length / x
        let step = parseInt((length / 100).toFixed(0))
        setChartData({
            options: {
                chart: {
                    id: "basic-bar"
                },
                xaxis: {
                    categories: weatherData.data.hourly.time.filter((_, idx) => idx % step === 0).map(isoTime => {
                        let parsedDate = new Date(isoTime)
                        let splitDate = parsedDate.toString().split(" ")
                        // Return the month & year
                        return `${splitDate[1]} ${splitDate[2]} ${splitDate[3]}`
                    })
                }
            },
            series: [
                {
                    name: "temps",
                    data: weatherData.data.hourly.temp.filter((_, idx) => idx % step === 0)
                }
            ]
        })
        setIsLoading(false)
    }

    return <Box>
        <Typography variant="h4" gutterBottom>Welcome back {auth.username}</Typography>
        <Typography variant="h5" gutterBottom>The default graph is your register hometown = "{userData.hometown}" for the last 3 months</Typography>
        {/*<Typography variant="h5" gutterBottom></Typography>*/}
        <ButtonGroup variant="contained" sx={{marginBottom: 3}} size="large">
            <Button name="3M"
                    onClick={() => fetchWeatherData(userData, "3M")}
            >3 Months</Button>
            <Button name="1Y"
                    onClick={() => fetchWeatherData(userData, "1Y")}
            >1 Year</Button>
            <Button name="5Y"
                    onClick={() => fetchWeatherData(userData, "5Y")}
            >5 Year</Button>
        </ButtonGroup>
        <Box display="flex" width="100%" justifyContent="center" alignItems="center" flexDirection="column">
            {!isLoading ? <Chart
                options={chartData.options}
                series={chartData.series}
                type="line"
                width="500"
            /> : <Box><CircularProgress/></Box>}
        </Box>
    </Box>
}

export default Dashboard