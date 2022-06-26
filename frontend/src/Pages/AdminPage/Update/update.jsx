import React, { Component } from "react";
import API from "../../../Api/Api";
import axios from "axios";
import { useParams } from "react-router-dom";

import { Container, Form, Button } from "react-bootstrap";

export class Update extends Component {
	// handleUpdate() {
	// 	console.log("update");
	// }

	state = {
		id: "",
		programming_languange: "",
		questions: "",
	};

	// componentDidMount = async () => {
	// 	const id = this.props.match.params.id;
	// 	const auth = localStorage.getItem("token");
	// 	const res = await axios.get(`${API.API_URL}/admin/questions/` + id, {
	// 		method: "GET",
	// 		headers: {
	// 			Accept: "application/json",
	// 			"Content-Type": "application/json",
	// 			Authorization: "Bearer " + auth,
	// 		},
	// 	});
	// 	this.setState({
	// 		id: res.data.id,
	// 		programming_languange: res.data.programming_languange,
	// 		questions: res.data.questions,
	// 	});
	// };

	// handlerChange = (e) => {
	// 	this.setState({ [e.target.name]: e.target.value });
	// };

	handleUpdate = async (e) => {
		const auth = localStorage.getItem("token");
		const id = this.props.match.params.id;
		e.preventDefault();
		await axios.put(`${API.API_URL}/admin/questions/update` + id, this.state);
	};
	render() {
		const { programming_languange, questions } = this.state;

		return (
			<section id="create-question">
				<Container className="container-outside">
					<Form onSubmit={this.handleUpdate}>
						<h1 className="mb-5">Update Question</h1>
						<Form.Group className="mb-3 mx-auto" controlId="formBasicEmail">
							<Form.Label className="card-title">
								Programming Language
							</Form.Label>
							<Form.Control
								type="text"
								className="mx-auto form-input"
								// value={programming_languange}
								onChange={this.handlerChange}
							/>
						</Form.Group>

						<Form.Group className="mb-3" controlId="formBasicPassword">
							<Form.Label className="card-title">Question</Form.Label>
							<Form.Control
								as="textarea"
								className="mx-auto form-area"
								// value={questions}
								onChange={this.handlerChange}
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
}

export default Update;

// import React, { useState } from "react";
// import axios from "axios";
// import "./update.css";
// import API from "../../../Api/Api";

// import { Container, Form, Button } from "react-bootstrap";

// export default function Update(id) {
// 	const apiQuestions = `${API.API_URL}/admin/questions/update/id`;
// 	let auth = localStorage.getItem("token");

// 	const [questions, setQuestions] = useState([]);
// 	const [programming_languange, setProgramming_languange] = useState("");

// 	const getPosts = async () => {
// 		const id = this.props.match.params.id;
// 		const res = await axios
// 			.get(apiQuestions + id, {
// 				headers: {
// 					Accept: "application/json",
// 					"Content-Type": "application/json",
// 					Authorization: "Bearer " + auth,
// 				},
// 			})
// 			.then((result) => {
// 				result.json().then((res) => {
// 					setQuestions(res.data.questions);
// 					setProgramming_languange(res.data.programming_languange);
// 					console.log(res.data);
// 				});
// 			});
// 		console.log(res.data);
// 	};

// 	getPosts();

// 	const handleSubmit = async (e) => {
// 		e.preventDefault();
// 		const id = this.props.match.params.id;
// 		const res = await axios
// 			.put(
// 				apiQuestions + `${id}`,
// 				{
// 					programming_languange: programming_languange,
// 				},
// 				{
// 					headers: {
// 						Accept: "application/json",
// 						"Content-Type": "application/json",
// 						Authorization: "Bearer " + auth,
// 					},
// 				}
// 			)
// 			.then(
// 				(result) => {
// 					result.json().then((res) => {
// 						console.log(res);
// 						getPosts();
// 					});
// 				}
// 				// console.log(res.data);
// 			);
// 		console.log(res.data);
// 	};

// 	return (
// 		<section id="create-question">
// 			<Container className="container-outside">
// 				<Form onSubmit={handleSubmit}>
// 					<h1 className="mb-5">Update Question</h1>
// 					<Form.Group className="mb-3 mx-auto" controlId="formBasicEmail">
// 						<Form.Label className="card-title">Programming Language</Form.Label>
// 						<Form.Control
// 							type="text"
// 							className="mx-auto form-input"
// 							value={programming_languange}
// 							onChange={(event) => setProgramming_languange(event.target.value)}
// 						/>
// 					</Form.Group>

// 					<Form.Group className="mb-3" controlId="formBasicPassword">
// 						<Form.Label className="card-title">Question</Form.Label>
// 						<Form.Control
// 							as="textarea"
// 							className="mx-auto form-area"
// 							value={questions}
// 							onChange={(event) => setQuestions(event.target.value)}
// 						/>
// 					</Form.Group>
// 					<Button variant="warning button" className="mt-3" type="submit">
// 						Update
// 					</Button>
// 				</Form>
// 			</Container>
// 		</section>
// 	);
// }
