import * as React from "react";
import "./style.css";
import banner from "../../Assets/Images/Login-Images.png";

const Login = () => {
	return (
		<section id="login-pages">
			<div className="row">
				<div className="login-left col-lg-5">
					<img src={banner} alt="banner" />
				</div>
				<div className="login-right col-lg-7 my-auto">
					<form>
						<h2 className="mb-5">Login to your Account</h2>
						<div className="mb-3">
							<label for="usernameInput" className="form-label">
								Username
							</label>
							<input
								type="email"
								className="form-control"
								id="usernameInput1"
								aria-describedby="usernameHelp"
							/>
							<div id="emailHelp" className="form-text">
								We'll never share your username with anyone else.
							</div>
						</div>
						<div className="mb-3">
							<label for="exampleInputPassword1" className="form-label">
								Password
							</label>
							<input
								type="password"
								className="form-control"
								id="exampleInputPassword1"
							/>
						</div>
						<div className="mb-3 form-check">
							<input
								type="checkbox"
								className="form-check-input"
								id="exampleCheck1"
							/>
							<label className="form-check-label" for="exampleCheck1">
								Check me out
							</label>
						</div>
						<button type="submit" className="btn mt-3">
							Submit
						</button>
					</form>
				</div>
			</div>
		</section>
	);
};

export default Login;