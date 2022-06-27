import React, { useState } from "react";
import axios from "axios";
import "./update.css";
import API from "../../../Api/Api";

import { Container, Form, Button } from "react-bootstrap";
import { useNavigate } from "react-router-dom";

export default function Update() {
	const navigate = useNavigate();
	const questionId = localStorage.getItem("questionId");
	const [programming_languange, setProgramming_languange] = useState("");
	const [question, setQuestion] = useState("");

	const handleUpdate = async (event) => {
		event.preventDefault();
		try {
			let res = await axios.put(
				`${API.API_URL}/admin/questions/update/` + questionId,
				{
					programming_languange: programming_languange,
					question: question,
				},
				{
					headers: {
						Accept: "/",
						"Content-Type": "application/json",
						Authorization: "Bearer " + localStorage.getItem("token"),
					},
				}
			);

			if (res.status === 200) {
				navigate("/Dashboard");
			}
		} catch (error) {}
	};

	return (
		<section id="create-question">
			<Container className="container-outside">
				<Form onSubmit={handleUpdate}>
					<h1 className="mb-5">Update Question</h1>
					<Form.Group className="mb-3 mx-auto" controlId="formBasicEmail">
						<Form.Label className="card-title">Programming Language</Form.Label>
						<Form.Control
							type="text"
							className="mx-auto form-input"
							value={programming_languange}
							onChange={(event) => setProgramming_languange(event.target.value)}
						/>
					</Form.Group>

					<Form.Group className="mb-3" controlId="formBasicPassword">
						<Form.Label className="card-title">Question</Form.Label>
						<Form.Control
							as="textarea"
							className="mx-auto form-area"
							value={question}
							onChange={(event) => setQuestion(event.target.value)}
						/>
					</Form.Group>
					<Button variant="warning button" className="mt-3" type="submit">
						Update
					</Button>
				</Form>
			</Container>
		</section>
	);
}
