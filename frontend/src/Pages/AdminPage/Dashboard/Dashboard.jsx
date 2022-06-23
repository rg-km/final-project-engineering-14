<<<<<<< HEAD
import React from "react";
import "./dashboard.css";
// import Banner from "../../Assets/Images/Level-image.png";

//React Bs
import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";

// import Image from "react-bootstrap/Image";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Navbar from "../../../Components/Navbar/Navbar";

export default function Dashboard() {
	return (
		<section id="dashboard-pages">
			<Navbar />
			<div className="profile-name pt-3">
				<h1>Welcome, Admin!</h1>
			</div>
			<div>
				<Container>
					<Card className="card-main">
						<Card.Body className="card-body">
							<Button className="button-1">Jumlah Soal : 10 </Button>
							<Button>New Question :</Button>
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
					</Card>
				</Container>
			</div>
		</section>
	);
}
=======
import React from "react";
import "./dashboard.css";
// import Banner from "../../Assets/Images/Level-image.png";

//React Bs
import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";

// import Image from "react-bootstrap/Image";
import { Table } from "react-bootstrap";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Navbar from "../../../Components/Navbar/Navbar";

export default function Dashboard() {
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
                <p className="my-auto">Jumlah Soal :</p>
              </Card>
            </Col>
            <Col className="float-sm-right">
              <Button
                variant="warning"
                className="button-create float-sm-right"
              >
                New Question
              </Button>
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
            <tbody className="table-body">
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
>>>>>>> 60e79c5a4a1ca2f7c740c9df40c294609c2df025
