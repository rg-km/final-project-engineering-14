import React from "react";
import "./recomendation.css";
import Navbar from "../../Components/Navbar/Navbar";
import Footer from "../../Components/Footer/Footer";
import axios from "axios";
import API from "../../Api/Api";

import { useNavigate } from "react-router-dom";
import { Container, Button } from "react-bootstrap";
import { useState, useEffect } from "react";

const Recomendation = () => {
	let x = localStorage.getItem("level_id");
	const navigate = useNavigate();

	const [recommendation, setRecommendation] = useState([]);

	const getRecommendation = async () => {
		let auth = localStorage.getItem("token");

		try {
			let res = await axios.get(`${API.API_URL}/home/recommendation/${x}`, {
				headers: {
					Accept: "/",
					"Content-Type": "application/json",
					Authorization: "Bearer " + auth,
				},
			});
			console.log(res);
			setRecommendation(res.data.data.image_url);
		} catch (error) {
			console.log(error);
		}
	};

	useEffect(() => {
		getRecommendation();
	}, []);

	const handleDone = () => {
		alert("Congratulations! You have finished the test");
		navigate("/HomePage");
	};
	return (
		<section id="recomendation-page">
			<Navbar />
			<h1 className="mb-5 text-center title">Recomendation Page</h1>
			<Container className="container-luar">
				<img src={recommendation} alt="recomendation" className="images" />
				<br />

				<Button variant="success" className="btn-done" onClick={handleDone}>
					SELESAI
				</Button>
			</Container>
			<Footer />
		</section>
	);
};

export default Recomendation;
