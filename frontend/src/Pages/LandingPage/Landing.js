import React from "react";
import Navbar from "./Navbar/navbar";
import Carousel from "./Carousel/Carousel";
import About from "./About/about";

const LandingPage = () => {
	return (
		<>
			<Navbar />
			<Carousel />
			<About />
		</>
	);
};

export default LandingPage;
