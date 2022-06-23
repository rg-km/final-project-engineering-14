import React, { useContext } from "react";
import AuthContext from "../../store/auth-context";
import { useNavigate } from "react-router-dom";

import "./navbar.css";
import Logo from "../../../src/Assets/Images/Logo.png";
// import Avatar from "../../../src/Assets/Images/avatar.png";

// React Boostrap
import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Button from "react-bootstrap/Button";
// import Card from "react-bootstrap/Card";
// import Image from "react-bootstrap/Image";

function HomeNavbar() {
	const authCtx = useContext(AuthContext);
	const navigate = useNavigate();

	const logoutHandler = () => {
		authCtx.logout();
		navigate("/");
	};

	return (
		<section id="navbar-home">
			<Navbar className="nav" expand="lg" variant="light">
				<Container>
					<Row>
						<Col>
							<img src={Logo} className="img-fluid" alt="logo" />
						</Col>
					</Row>
					<Row>
						<Navbar.Toggle aria-controls="basic-navbar-nav" />
						<Navbar.Collapse id="basic-navbar-nav" className="nav-content mx-3">
							<Nav className="me-auto">
								<Nav.Link href="#home" className="text-light">
									Home
								</Nav.Link>
								<Nav.Link href="#link" className="text-light">
									Journey
								</Nav.Link>
								<Nav.Link href="#link" className="text-light">
									Learning Path
								</Nav.Link>
							</Nav>
							<Nav className="text-light">
								<Button variant="warning" onClick={logoutHandler}>
									Logout
								</Button>
							</Nav>
							{/* <div className="card-profile">
								<Card style={{ width: "200px", height: "60px" }}>
									<Row>
										<Col>
											<Image
												className="card-avatar"
												src={Avatar}
												roundedCircle
											/>
										</Col>
										<Col className="col-8 my-auto">
											<h5>Rahman Budi</h5>
										</Col>
									</Row>
								</Card>
							</div> */}
						</Navbar.Collapse>
					</Row>
				</Container>
			</Navbar>
		</section>
	);
}

export default HomeNavbar;
