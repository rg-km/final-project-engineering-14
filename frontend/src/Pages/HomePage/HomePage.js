import React from "react";
import Main from "./Main/Main";
import Navbar from "../../Components/Navbar/Navbar";
import LearningPath from "./LearningPath/LearningPath";

function HomePage() {
	return (
		<>
			<Navbar />
			<Main />
			<LearningPath />
		</>
	);
}

export default HomePage;
