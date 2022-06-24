import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

import "./createQuestion.css";
import Navbar from "../../../Components/Navbar/Navbar";

// React Boostraps
import { Container, Form, Button } from "react-bootstrap";

export default function CreateQuestion() {
	// const navigate = useNavigate();
	const [programming_languange, setProgramming_languange] = useState("");
	const [question, setQuestion] = useState("");
	const navigate = useNavigate();

	const handleCreate = async (event) => {
		event.preventDefault();
		try {
			let res = await axios.post(
				"https://33d9-120-188-66-111.ap.ngrok.io/admin/questions/create",
				{
					programming_languange: programming_languange,
					question: question,
					returnSecureToken: true,
				},
				{
					headers: {
						Accept: "/",
						"Content-Type": "application/json",
						Authorization: "Bearer " + localStorage.getItem("token"),
					},
				}
			);
			// console.log(res);
			if (res.status === 200) {
				navigate("/Dashboard");
			}
		} catch (error) {
			// alert(
			// 	"Username / Email Sudah terdaftar, Silahkan Periksa Data Anda Kembali!"
			// );
			console.log(error);
		}
	};

	return (
		<section id="create-question">
			<Navbar />

			<Container className="container-outside">
				<Form onSubmit={handleCreate}>
					<h1 className="mb-5">Create Question</h1>
					<Form.Group className="mb-3 mx-auto" controlId="formBasicEmail">
						<Form.Label className="card-title">Programming Language</Form.Label>
						<Form.Control
							type="text"
							className="mx-auto form-input"
							onChange={(event) => setProgramming_languange(event.target.value)}
						/>
					</Form.Group>

					<Form.Group className="mb-3" controlId="formBasicPassword">
						<Form.Label className="card-title">Question</Form.Label>
						<Form.Control
							as="textarea"
							className="mx-auto form-area"
							onChange={(event) => setQuestion(event.target.value)}
						/>
					</Form.Group>
					<Button variant="warning button" className="mt-3" type="submit">
						Submit
					</Button>
				</Form>
			</Container>
		</section>
	);
}
