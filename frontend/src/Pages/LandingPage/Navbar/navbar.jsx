import React from "react";
import Logo from "../../../Assets/Images/Logo.png";
import "./style.css";

const navbar = () => {
	return (
		<section id="navbar-landing">
			<div className="container">
				<div className="row">
					<div className="left-content col-4">
						<img src={Logo} className="img-fluid" alt="logo" />
					</div>
					<div className="right-content col-8 my-auto ">
						<button className="signup float-lg-right">Sign Up</button>
						<button className="login float-lg-right mr-2">Log In</button>
					</div>
				</div>
			</div>
		</section>
	);
};

export default navbar;
