import React from "react";
import { Link } from "react-router-dom";
import Logo from "../../../Assets/Images/Logo.png";
import "./style.css";

// React Boostrap
import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";

const navbar = () => {
  return (
    <section id="navbar-landing">
      <Navbar className="nav" expand="lg" variant="light">
        <Container>
          <Row>
            <Col>
              <img src={Logo} className="img-fluid" alt="logo" />
            </Col>
          </Row>
          <Row className="right-content">
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav">
              <Nav className="me-auto">
                <Nav.Link>
                  <Link to="/register">
                    <button className="signup float-lg-right">Sign Up</button>
                  </Link>
                </Nav.Link>
                <Nav.Link>
                  <Link to="/login">
                    <button className="login float-lg-right mr-2">
                      Log In
                    </button>
                  </Link>
                </Nav.Link>
              </Nav>
            </Navbar.Collapse>
          </Row>
        </Container>
      </Navbar>
    </section>
  );
};

export default navbar;
