import React, { useState, useEffect } from "react";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";

import "./dashboard.css";
import Navbar from "../../../Components/Navbar/Navbar";
import { Container, Table, Card, Button, Row, Col } from "react-bootstrap";

export default function Dashboard() {
	const navigate = useNavigate();
	const [posts, setPosts] = useState([]);
	const [check, setCheck] = useState(false);
	const apiQuestions =
		"https://33d9-120-188-66-111.ap.ngrok.io/admin/questions";
	const apiDeletes =
		"https://33d9-120-188-66-111.ap.ngrok.io/admin/questions/delete/";

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
										<Link to={"/Update/" + post.id}>
											<Button
												variant="warning"
												// onClick={() => handleUpdate(questionId)}
											>
												Update
											</Button>
										</Link>
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

// const [dashboard, setDashboard] = useState([]);
// const [programming_languange, setProgrammingLanguange] = useState("");
// const [question, setQuestion] = useState("");
// const [id, setQuestionId] = useState(null);

// useEffect(() => {
// 	getUser();
// }, []);

// function getUser() {
// 	let auth = localStorage.getItem("token");
// 	fetch("https://db77-120-188-37-170.ap.ngrok.io/admin/questions", {
// 		headers: {
// 			Accept: "application/json",
// 			"Content-Type": "application/json",
// 			Authorization: "Bearer " + auth,
// 		},
// 	}).then((result) => {
// 		result.json().then((res) => {
// 			setDashboard(res);
// 			setProgrammingLanguange(res[0].programming_languange);
// 			setQuestion(res[0].question);
// 			setQuestionId(res[0].id);
// 		});
// 	});
// }
// function deleteUser(id) {
// 	fetch("https://db77-120-188-37-170.ap.ngrok.io", {
// 		method: "DELETE",
// 	}).then((result) => {
// 		result.json().then((res) => {
// 			console.warn(res);
// 			getUser();
// 		});
// 	});
// }
// function selectUser(id) {
// 	let item = users[id - 1];
// 	setProgramming_Languange(item.programming_languange);
// 	setQuestion(item.question);
// }

// const getData = () => {
// 	let auth = localStorage.getItem("token");
// 	fetch("https://db77-120-188-37-170.ap.ngrok.io/admin/questions", {
// 		method: "GET",
// 		headers: {
// 			Accept: "application/json",
// 			"Content-Type": "application/json",
// 			Authorization: "Bearer " + auth,
// 		},
// 	})
// 		// .then((res) => res.json());
// 		.then((data) => {
// 			return data.json();
// 		})
// 		.then((objectData) => {
// 			console.log(objectData.data);
// 			let tableData = "";
// 			objectData.data.map((values) => {
// 				tableData += `
//       <tr>
//         <td>${values.programming_languange}</td>
//         <td>${values.question}</td>
//         <td><Button class="update">Update</Button></td>
//         <td><Button class="delete">Delete</Button></td>
//       </tr>
//       `;
// 			});
// 			document.getElementById("table-body").innerHTML = tableData;
// 		})
// 		.catch((err) => {
// 			console.log(err);
// 		});
// };

// getData();

// const getDataTotal = () => {
// 	fetch("https://9576-114-4-212-203.ap.ngrok.io/admin/dashboard", {
// 		method: "GET",
// 		headers: {
// 			Accept: "application/json",
// 			"Content-Type": "application/json",
// 			Authorization: "Bearer " + auth,
// 		},
// 	})
// 		// .then((res) => res.json());
// 		.then((data) => {
// 			return data.json();
// 		})
// 		.then((objectData) => {
// 			console.log(objectData.data[0].total_question);
// 			// let tableData = "";
// 			// objectData.data.map((values) => {
// 			// 	tableData = `<p>${values.total_question}</p>`;
// 			// });
// 			// document.getElementById("total_question").innerHTML = tableData;
// 		})
// 		.catch((err) => {
// 			console.log(err);
// 		});
// };

// getDataTotal();

// function buat fetch data all //

// const delete_data = () => {
// 	const deleteData = () => {
// 		let auth = localStorage.getItem("token");
// 		fetch(
// 			"https://3090-103-169-136-72.ap.ngrok.io/admin/questions/delete/{questionId}",
// 			{
// 				method: "Delete",
// 				headers: {
// 					Accept: "application/json",
// 					"Content-Type": "application/json",
// 					Authorization: "Bearer " + auth,
// 				},
// 			}
// 		)
// 			// .then((res) => res.json());
// 			// .then((res) => {
// 			//   if(res.ok) => {
// 			//               console.log("Delete Success");
// 			//           }
// 			// })
// 			.catch((err) => {
// 				console.log(err);
// 			});
// 	};
// 	return deleteData();
// };
