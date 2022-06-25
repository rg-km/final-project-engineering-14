import React, { useState, useEffect } from "react";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";
import API from "../../../Api/Api";

import Navbar from "../../../Components/Navbar/Navbar";
import { Container, Form, Button } from "react-bootstrap";

export default function Update(props) {
	const [programming_languange, setProgramming_languange] = useState("");
	const [question, setQuestion] = useState("");
	const navigate = useNavigate();
	const [posts, setPosts] = useState([]);
	const [check, setCheck] = useState(false);
	const apiQuestions = `${API.API_URL}/admin/questions`;

	let auth = localStorage.getItem("token");

	useEffect(() => {
		getPosts();
	}, []);

	const getPosts = async () => {
		const id = props.match.params.id;
		const { data: res } = await axios.get(apiQuestions + id, {
			headers: {
				Accept: "application/json",
				"Content-Type": "application/json",
				Authorization: "Bearer " + auth,
			},
		});
		console.log(res.data);
		// console.log(posts.data.length);
		// if (!res.data === "null") {
		setPosts(res.data);
		setCheck(true);
		// } else {
		// 	setPosts(res.data);
		// 	setCheck(false);
		// }
	};
	const handleCreate = async (event) => {
		event.preventDefault();
		try {
			let res = await axios.post(
				"https://8b7d-120-188-34-21.ap.ngrok.io/admin/questions/create",
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
		<section id="dashboard-pages">
			<Navbar />
			<div className="profile-name pt-3">
				<h1 className="mb-5">Welcome, Admin!</h1>
			</div>
			<div>
				<Container className="container-outside">
					<Form onSubmit={handleCreate}>
						<h1 className="mb-5">Create Question</h1>
						<Form.Group className="mb-3 mx-auto" controlId="formBasicEmail">
							<Form.Label className="card-title">
								Programming Language
							</Form.Label>
							<Form.Control
								type="text"
								className="mx-auto form-input"
								onChange={(event) =>
									setProgramming_languange(event.target.value)
								}
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
			</div>
		</section>
	);
}

// import React, { useState } from "react";
// import axios from "axios";
// import "./update.css";

// import { Container, Form, Button } from "react-bootstrap";

// export default function Update() {
// 	const apiQuestions =
// 		"https://33d9-120-188-66-111.ap.ngrok.io/admin/questions/update/";
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
// 					<h1 className="mb-5">Create Question</h1>
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
// 						Submit
// 					</Button>
// 				</Form>
// 			</Container>
// 		</section>
// 	);
// }
