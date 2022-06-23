// import * as React from "react";
import React, { useRef, useContext } from "react";
import { useNavigate } from "react-router-dom";
import AuthContext from "../../store/auth-context";

import "./style.css";
import banner from "../../Assets/Images/Login-Images.png";

export const Login = () => {
	const navigate = useNavigate();

	const emailInputRef = useRef();
	const passwordInputRef = useRef();

	const authCtx = useContext(AuthContext);

	if (authCtx.role === "admin") {
		navigate("/Dashboard");
	} else if (authCtx.isLoggedIn) {
		navigate("/HomePage");
	}

	const handleSubmit = (event) => {
		event.preventDefault();

		const enteredEmail = emailInputRef.current.value;
		const enteredPassword = passwordInputRef.current.value;

		let url;
		url = "https://d376-120-188-37-170.ap.ngrok.io/auth/sign-in";
		fetch(url, {
			method: "POST",
			body: JSON.stringify({
				email: enteredEmail,
				password: enteredPassword,
				returnSecureToken: true,
			}),
			headers: {
				Accept: "/",
				"Content-Type": "application/json",
			},
		})
			.then((res) => {
				if (res.ok) {
					return res.json();
				} else {
					return res.json().then((data) => {
						let errorMassage = "Authentication failed!";
						// if (data && data.error && data.error.message) {
						// 	errorMassage = data.error.message;
						// }

						throw new Error(errorMassage);
					});
				}
			})
			.then((data) => {
				// console.log(data.data.token);
				authCtx.login(data.data.token, data.data.role);
			})
			.catch((err) => {
				// alert(err.message);
				alert("Data yang anda masukkan salah, silahkan di cek!");
			});
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
					<form onSubmit={handleSubmit}>
						<div className="mb-3">
							<label for="emailInput" className="form-label">
								Email
							</label>
							<input
								type="email"
								id="email"
								className="form-control"
								ref={emailInputRef}
							/>
						</div>
						<div className="mb-3">
							<label for="exampleInputPassword1" className="form-label">
								Password
							</label>
							<input
								type="password"
								id="password"
								className="form-control"
								ref={passwordInputRef}
							/>
						</div>
						<button type="submit" className="btn mt-3">
							Login
						</button>
					</form>
				</div>
			</div>
		</section>
	);
};

export default Login;
