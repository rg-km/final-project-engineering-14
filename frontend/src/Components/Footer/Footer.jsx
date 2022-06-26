import React from "react";
import { Link } from "react-router-dom";
import "./footer.css";
import Logo from "../../../src/Assets/Images/Logo.png";

//React Icon
import {
  FaFacebookSquare,
  FaYoutube,
  FaInstagram,
  FaTwitterSquare,
} from "react-icons/fa";

// React Boostrap
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";

const Footer = () => (
  <section id="footer">
    <Container className="main-content  text-light">
      <Row>
        <Col>
          <img src={Logo} className="img-fluid" alt="logo" />
        </Col>
        <Col className="pt-5">
          <h5>About us</h5>
        </Col>
        <Col className="pt-5">
          <h5>Follow us on</h5>
        </Col>
      </Row>
      <Row>
        <Col>
          <div className="company-name">
            <div>RUANGARAH HQ</div>
            <div className="pt-2">
              Jl. Tebet Barat Raya, Tebet Barat <br></br>
              Tebet, Kota Jakarta Selatan, <br></br>
              Daerah Khusus Ibukota Jakarta 12820
            </div>
          </div>
        </Col>
        <Col>
          <div className="about-company">
            <div className="pt-2">
              <Link to="/" className="pages-link">
                About
              </Link>
            </div>
            <div className="pt-2">
              <Link to="/" className="pages-link">
                Contact Us
              </Link>
            </div>
            <div className="pt-2">
              <Link to="/" className="pages-link">
                Out Teams
              </Link>
            </div>
          </div>
        </Col>
        <Col>
          <div className="follow-us">
            <div className="double-logo">
              <div className="facebook">
                <Link to="/" className="pages-link">
                  <FaFacebookSquare />
                  <span className="px-2">Facebook</span>
                </Link>
              </div>
              <div className="youtube mx-3">
                <Link to="/" className="pages-link">
                  <FaYoutube />
                  <span className="px-2">YouTube</span>
                </Link>
              </div>
            </div>
            <div className="pt-2">
              <Link to="/" className="pages-link">
                <FaInstagram />
                <span className="px-2">Instagram</span>
              </Link>
            </div>
            <div className="pt-2">
              <Link to="/" className="pages-link">
                <FaTwitterSquare />
                <span className="px-2">Twitter</span>
              </Link>
            </div>
          </div>
        </Col>
      </Row>
      <hr className="footer-line text-light" />
      <Row>
        <Col>
          <div className="copyright">
            <h5>&copy; 2022 PT RUANGARAH INDONESIA. All Rights Reserved.</h5>
          </div>
        </Col>
      </Row>
    </Container>
  </section>
);

export default Footer;
