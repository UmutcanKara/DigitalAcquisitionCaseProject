import React from "react";
import {Box, Container, Button, FormHelperText, FormControl, TextField, Typography} from "@mui/material";
import {AuthService} from "../../api/service/auth.js";
import {AuthContext} from "../../context/AuthContext.jsx";
import {useNavigate} from "react-router-dom";

const Login = () => {
    const [formData, setFormData] = React.useState({
        username: "",
        password: ""
    })
    const navigate = useNavigate()
    const { login } = React.useContext(AuthContext)
    const onChangeHandler = (e) => {
        setFormData(prevState => ({
            ...prevState,
            [e.target.name]: e.target.value,
        }))
    }
    const onSubmitHandler = async () => {
        let response = await AuthService.login(formData)
        if (response.status === 200) {
            login(response.data.token, formData.username)
            navigate("/dashboard")
        }
    }

    return <Box marginBottom={3}>
        <Typography variant="h4" gutterBottom>Login</Typography>
        <Box gap={2} display="flex" flexDirection="column" >
            <TextField value={formData.username} onChange={onChangeHandler} name="username" label="Username" required />
            <TextField value={formData.password} onChange={onChangeHandler} name="password" label="Password" type="password" required />
            <Button variant="contained" sx={{paddingY: 1}} onClick={onSubmitHandler} >Login</Button>
        </Box>
    </Box>
}

export default Login