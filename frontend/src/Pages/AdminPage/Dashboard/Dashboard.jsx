import React, { useState, useEffect } from "react";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";
import API from "../../../Api/Api";

import "./dashboard.css";
import Navbar from "../../../Components/Navbar/Navbar";
import { Container, Table, Card, Button, Row, Col } from "react-bootstrap";

export default function Dashboard(props) {
	const [posts, setPosts] = useState([]);
	const [check, setCheck] = useState(false);
	const [edit, setEdit] = useState("");

	const navigate = useNavigate();

	const apiQuestions = `${API.API_URL}/admin/questions`;
	const apiDeletes = `${API.API_URL}/admin/questions/delete/`;

	let auth = localStorage.getItem("token");

	useEffect(() => {
		getPosts();
	}, []);

	const getPosts = async () => {
		const { data: res } = await axios.get(apiQuestions, {
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

	console.warn(posts);

	function deleteQuestion(id) {
		fetch(apiDeletes + `${id}`, {
			method: "DELETE",
			headers: {
				Accept: "application/json",
				"Content-Type": "application/json",
				Authorization: "Bearer " + auth,
			},
		})
			// console.log(posts.questionQuestion);
			.then((result) => {
				result.json().then((res) => {
					console.warn(res);
					getPosts();
				});
			});
	}

	function handleUpdate(id) {
		console.log(id);
		// navigate.push("/Update" + `${id}`);
	}

	return (
		<section id="dashboard-pages">
			<Navbar />
			<div className="profile-name pt-3">
				<h1 className="mb-5">Welcome, Admin!</h1>
			</div>
			<div>
				<Container>
					<Row className="mb-5">
						<Col>
							{/* <Card className="card float-left text-center">
								<Row className="my-auto mt-2">
									<Col className="col-9">
										<p className="my-auto">Jumlah Soal : {posts.length}</p>
									</Col>
									<Col className="col-3 ps-1"> */}
							{/* {posts.map((postlength, index) => {
											<p key={index} id="total_question">
												{postlength.length}
											</p>;
										})} */}
							{/* </Col>
								</Row>
							</Card> */}
						</Col>
						<Col className="float-sm-right">
							<Link to="/CreateQuestion">
								<Button
									variant="warning"
									className="button-create float-sm-right"
								>
									New Question
								</Button>
							</Link>
						</Col>
					</Row>
					<Table className="table">
						<thead>
							<tr>
								<th>Programming Language</th>
								<th>Question</th>
								<th>Update</th>
								<th>Delete</th>
							</tr>
						</thead>
						<tbody id="table-body" className="table-body">
							{posts.map((post) => (
								<tr key={post.id}>
									<td className="my-auto">{post.programming_languange}</td>
									<td>{post.question}</td>
									<td>
										{/* <Link to={"/Update/" + post.id}> */}
										<Button
											variant="warning"
											onClick={() => handleUpdate(post.id)}
										>
											Update
										</Button>
										{/* </Link> */}
									</td>
									<td>
										<Button
											variant="danger"
											// onClick={handleDelete}
											onClick={() => deleteQuestion(post.id)}
										>
											Delete
										</Button>
									</td>
								</tr>
							))}
						</tbody>
					</Table>
				</Container>
			</div>
		</section>
	);
}
