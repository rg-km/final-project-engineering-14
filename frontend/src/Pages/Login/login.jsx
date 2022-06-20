// import * as React from "react";
import { Link } from "react-router-dom";
import { useState, React } from "react";
import axios from "axios";
import "./style.css";
import banner from "../../Assets/Images/Login-Images.png";

const Login = () => {
	const [email, setemail] = useState("");
	const [password, setpassword] = useState("");

	const handleSubmit = async (e) => {
		e.preventDefault();
		try {
			const response = await axios.post(
				"https://f61a-120-188-92-248.ap.ngrok.io/auth/sign-in",
				{
					email: email,
					password: password,
				},
				{
					headers: {
						Accept: "/",
						"User-Agent": "Thunder Client (https://www.thunderclient.com/)",
						"Content-Type": "application/json",
					},
				}
			);
			console.log(response);
		} catch (error) {
			console.log(error);
		}

		// axios
		// 	.post(
		// 		`https://7dbd-114-10-64-47.ap.ngrok.io/auth/sign-in`,
		// 		{
		// 			email: email,
		// 			password: password,
		// 		},
		// 		{
		// 			headers: {
		// 				Accept: "/",
		// 				"User-Agent": "Thunder Client (https://www.thunderclient.com/)",
		// 				"Content-Type": "application/json",
		// 			},
		// 		}
		// 	)
		// 	.then(function (response) {
		// 		// console.log(username, password);
		// 	});
	};

	return (
		<section id="login-pages">
			<div className="row">
				<div className="login-left col-lg-6">
					<img className="w-75 img-fluid mx-auto" src={banner} alt="banner" />
				</div>
				<div className="login-right col-lg-6 my-auto">
					<p>Welcome Back</p>
					<h2 className="mb-5">Login to your Account</h2>
					<form>
						<div className="mb-3">
							<label for="emailInput" className="form-label">
								Email
							</label>
							<input
								type="email"
								value={email}
								className="form-control"
								onChange={(e) => setemail(e.target.value)}
							/>
						</div>
						<div className="mb-3">
							<label for="exampleInputPassword1" className="form-label">
								Password
							</label>
							<input
								type="password"
								value={password}
								className="form-control"
								onChange={(e) => setpassword(e.target.value)}
							/>
						</div>
						{/* <div className="mb-3 form-check">
							<input
								type="checkbox"
								className="form-check-input"
								id="exampleCheck1"
							/>
							<label className="form-check-label" for="exampleCheck1">
								Check me out
							</label>
						</div> */}
						{/* <Link to="/home"> */}
						<button type="submit" className="btn mt-3" onClick={handleSubmit}>
							Login
						</button>
						{/* </Link> */}
					</form>
				</div>
			</div>
		</section>
	);
};

export default Login;
