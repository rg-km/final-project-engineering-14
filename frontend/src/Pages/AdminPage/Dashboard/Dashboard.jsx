import React, { useState, useEffect } from "react";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";
import API from "../../../Api/Api";

import "./dashboard.css";
import Navbar from "../../../Components/Navbar/Navbar";
import { Container, Table, Button, Row, Col, Card } from "react-bootstrap";

export default function Dashboard(id) {
	const [posts, setPosts] = useState([]);
	const navigate = useNavigate();

	const apiQuestions = `${API.API_URL}/admin/questions`;
	const apiDeletes = `${API.API_URL}/admin/questions/delete/`;

	let auth = localStorage.getItem("token");

	const getPosts = async () => {
		const { data: res } = await axios.get(apiQuestions, {
			headers: {
				Accept: "application/json",
				"Content-Type": "application/json",
				Authorization: "Bearer " + auth,
			},
		});
		setPosts(res.data);
	};

	useEffect(() => {
		getPosts();
	}, []);

	function deleteQuestion(id) {
		fetch(apiDeletes + `${id}`, {
			method: "DELETE",
			headers: {
				Accept: "application/json",
				"Content-Type": "application/json",
				Authorization: "Bearer " + auth,
			},
		}).then((result) => {
			result.json().then((res) => {
				getPosts();
			});
		});
	}

	function handleUpdate(id) {
		localStorage.setItem("questionId", id);
		navigate("/Update");
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
							<Card className="card float-left text-center">
								<p className="my-auto mx-auto">Jumlah Soal : {posts.length}</p>
							</Card>
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
								<tr className="tr-class" key={post.id}>
									<td className="my-auto">{post.programming_languange}</td>
									<td>{post.question}</td>
									<td>
										<Button
											variant="warning"
											onClick={() => handleUpdate(post.id)}
										>
											Update
										</Button>
									</td>
									<td>
										<Button
											variant="danger"
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
