import React, { useState } from "react";
import axios from "axios";
import "./update.css";

import { Container, Form, Button } from "react-bootstrap";

export default function Update() {
	const apiQuestions =
		"https://33d9-120-188-66-111.ap.ngrok.io/admin/questions/update/";
	let auth = localStorage.getItem("token");

	const [questions, setQuestions] = useState([]);
	const [programming_languange, setProgramming_languange] = useState("");

	const getPosts = async () => {
		const id = this.props.match.params.id;
		const res = await axios
			.get(apiQuestions + `${id}`, {
				headers: {
					Accept: "application/json",
					"Content-Type": "application/json",
					Authorization: "Bearer " + auth,
				},
			})
			.then((result) => {
				result.json().then((res) => {
					setQuestions(res.data.questions);
					setProgramming_languange(res.data.programming_languange);
					console.log(res.data);
				});
			});
		console.log(res.data);
	};

	getPosts();

	const handleSubmit = async (e) => {
		e.preventDefault();
		const id = this.props.match.params.id;
		const res = await axios
			.put(
				apiQuestions + `${id}`,
				{
					programming_languange: programming_languange,
				},
				{
					headers: {
						Accept: "application/json",
						"Content-Type": "application/json",
						Authorization: "Bearer " + auth,
					},
				}
			)
			.then(
				(result) => {
					result.json().then((res) => {
						console.log(res);
						getPosts();
					});
				}
				// console.log(res.data);
			);
		console.log(res.data);
	};

	return (
		<section id="create-question">
			<Container className="container-outside">
				<Form onSubmit={"#"}>
					<h1 className="mb-5">Create Question</h1>
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
							value={questions}
							onChange={(event) => setQuestions(event.target.value)}
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
