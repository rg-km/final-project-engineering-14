import React from "react";
import Logo from "../../../src/Assets/Images/Logo.png";
import Avatar from "../../../src/Assets/Images/avatar.png";
import "./navbar.css";

// React Boostrap
import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Card from "react-bootstrap/Card";
import Image from "react-bootstrap/Image";

function HomeNavbar() {
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
              <div className="card-profile">
                <Card style={{ width: "200px", height: "60px" }}>
                  <Image className="card-avatar" src={Avatar} roundedCircle />
                </Card>
              </div>
            </Navbar.Collapse>
          </Row>
        </Container>
      </Navbar>
    </section>
  );
}

export default HomeNavbar;
