import React from "react";
import { Container, Grid, Typography, Button } from "@mui/material";
import banner from "../../Assets/Images/Login-Images.png";

const Login = () => {
	return (
		<div>
			<Container>
				<Grid container spacing={4}>
					<Grid item xs={12} sm={6}>
						<Typography className="tittle" variant="h2">
							Everything You Need is{" "}
							<span className="caption">Already Inside</span>
						</Typography>
						<Button className="shopping-button">Get Product</Button>
					</Grid>

					<Grid className="brand" item sm={6}>
						<img src={banner} alt="banner" />
					</Grid>
				</Grid>
			</Container>
		</div>
	);
};

export default Login;
