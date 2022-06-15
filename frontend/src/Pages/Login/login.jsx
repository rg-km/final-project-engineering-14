import * as React from "react";
import "./style.css";
import banner from "../../Assets/Images/Login-Images.png";

const Login = () => {
	return (
		<section id="login-pages">
			<div className="row">
				<div className="test1 col-4">
					<img src={banner} alt="banner" />
				</div>
				<div className="test2 col-8">
					<button className="btn-danger">test</button>
				</div>
			</div>
		</section>
	);
};

export default Login;
