import React from "react";
import Logo from "../../Assets/Images/Logo.png";

const Navbar = () => {
	return (
		<section id="navbar">
			<nav class="navbar navbar-expand-lg">
				<div class="container-fluid">
					<img src={Logo} alt="logo" />
				</div>
			</nav>
			{/* <div className="container">
				<div className="row">
					<div className="col-4">
						<img src={Logo} alt="logo" />
					</div>
					<div className="col-8"></div>
				</div>
			</div> */}
		</section>
	);
};

export default Navbar;
