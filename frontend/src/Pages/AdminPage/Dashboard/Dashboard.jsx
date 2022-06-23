import React, { useState } from "react";
import axios from "axios";
import { Link } from "react-router-dom";

import "./dashboard.css";
import Navbar from "../../../Components/Navbar/Navbar";
// import Banner from "../../Assets/Images/Level-image.png";

//React Bs
import { Container, Table, Card, Button, Row, Col } from "react-bootstrap";

export default function Dashboard() {
	const getData = () => {
		let auth = localStorage.getItem("token");
		fetch("https://3090-103-169-136-72.ap.ngrok.io/admin/questions", {
			method: "GET",
			headers: {
				Accept: "application/json",
				"Content-Type": "application/json",
				Authorization: "Bearer " + auth,
			},
		})
			// .then((res) => res.json());
			.then((data) => {
				return data.json();
			})
			.then((objectData) => {
				// console.log(objectData.data.question);
				let tableData = "";
				objectData.data.map((values) => {
					tableData += `
          <tr>
            <td>${values.programming_languange}</td>
            <td>${values.question}</td>
            <td><Button>Update</Button></td>
            <td><Button>Delete</Button></td>
          </tr>
          `;
				});
				document.getElementById("table-body").innerHTML = tableData;
			})
			.catch((err) => {
				console.log(err);
			});
	};

	getData();

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
								<p className="my-auto">Jumlah Soal : </p>
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
							<tr>
								<td className="my-auto">Mark</td>
								<td>Otto</td>
								<td>
									<Button variant="warning">Update</Button>
								</td>
								<td>
									<Button variant="danger">Delete</Button>
								</td>
							</tr>
						</tbody>
					</Table>
					{/* <Card className="card-main">
            <Card.Body className="card-body">
              <Card>Jumlah Soal</Card>
              <Button>New Question</Button>
              <div className="card-question mt-5">
                <Row>
                  <Col sm={5}>
                    <div className="question-text">
                      <p>
                        Lorem Ipsum is simply dummy text of the printing and
                        typesetting industry. Lorem Ipsum has been the
                        industry's standard dummy text ever since the 1500s,
                        when an unknown printer took a galley of type and
                        scrambled it to make a type specimen book.
                      </p>
                    </div>
                    <div className="question-text">
                      <p>
                        Lorem Ipsum is simply dummy text of the printing and
                        typesetting industry. Lorem Ipsum has been the
                        industry's standard dummy text ever since the 1500s,
                        when an unknown printer took a galley of type and
                        scrambled it to make a type specimen book.
                      </p>
                    </div>
                    <div className="question-text">
                      <p>
                        Lorem Ipsum is simply dummy text of the printing and
                        typesetting industry. Lorem Ipsum has been the
                        industry's standard dummy text ever since the 1500s,
                        when an unknown printer took a galley of type and
                        scrambled it to make a type specimen book.
                      </p>
                    </div>
                  </Col>
                  <Col sm={2}>
                    <div className="button-update">
                      <Button>Update</Button>
                    </div>
                    <div className="button-update2 mt-5">
                      <Button>Update</Button>
                    </div>
                    <div className="button-update2 mt-5">
                      <Button>Update</Button>
                    </div>
                  </Col>
                  <Col sm={2}>
                    <div className="button-delete">
                      <Button>Delete</Button>
                    </div>
                    <div className="button-delete2 mt-5">
                      <Button>Delete</Button>
                    </div>
                    <div className="button-delete2 mt-5">
                      <Button>Delete</Button>
                    </div>
                  </Col>
                </Row>
              </div>
            </Card.Body>
          </Card> */}
				</Container>
			</div>
		</section>
	);
}
