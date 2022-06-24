import React from "react";
import "./recomendation.css";
import Navbar from "../../Components/Navbar/Navbar";
import Footer from "../../Components/Footer/Footer";

import { Container } from "react-bootstrap";

const recomendation = () => {
	return (
		<section id="recomendation-page">
			<Navbar />
			<h1 className="mb-5 text-center title">Recomendation Page</h1>
			<Container className="container-luar"></Container>
			<Footer />
		</section>
	);
};

export default recomendation;
