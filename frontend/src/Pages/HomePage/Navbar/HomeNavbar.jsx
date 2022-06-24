import React, { useContext } from "react";
import AuthContext from "../../../store/auth-context";

// Style React Boostrap
import "./style.css";
import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
// import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/esm/Button";

// Assets
import Logo from "../../../Assets/Images/Logo.png";
import { useNavigate } from "react-router-dom";
// import Avatar from "../../../Assets/Images/avatar.png";

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
								<Nav.Link href="" className="text-light pe-3">
									Home
								</Nav.Link>
								<Nav.Link href="#link" className="text-light pe-4">
									Learning Path
								</Nav.Link>
								<Nav className="text-light">
									<Button variant="warning" onClick={logoutHandler}>
										Logout
									</Button>
								</Nav>
							</Nav>
							{/* <Card style={{ width: "200px", height: "60px" }}>
								<Card.Img src={Avatar} className="card-avatar" /> */}

							{/* <Card.Text>Your Name</Card.Text> */}
							{/* </Card> */}
						</Navbar.Collapse>
					</Row>
				</Container>
			</Navbar>
		</section>
	);
}

export default HomeNavbar;