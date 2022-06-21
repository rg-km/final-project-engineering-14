import React from "react";
import Main from "./Main/Main";
import HomeNavbar from "./Navbar/HomeNavbar";
import LearningPath from "./LearningPath/LearningPath";

function HomePage() {
	return (
		<>
			<HomeNavbar />
			<Main />
			<LearningPath />
		</>
	);
}

export default HomePage;
