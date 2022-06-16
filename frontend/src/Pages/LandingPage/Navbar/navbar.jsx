import React from "react";
import { Link } from "react-router-dom";
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
						<Link to="/register">
							<button className="signup float-lg-right">Sign Up</button>
						</Link>
						<Link to="/login">
							<button className="login float-lg-right mr-2">Log In</button>
						</Link>
					</div>
				</div>
			</div>
		</section>
	);
};

export default navbar;
