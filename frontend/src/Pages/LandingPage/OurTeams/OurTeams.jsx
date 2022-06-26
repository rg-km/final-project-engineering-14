import React from "react";
import "./OurTeams.css";
import profile from "../../../Assets/Images/avatar.png";
import Footer from "../../../Components/Footer/Footer";

// React BS
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Card from "react-bootstrap/Card";
import Container from "react-bootstrap/Container";

export default function OurTeams() {
  return (
    <section id="OurTeams-page">
      <Container>
        <div className="text-heading">
          <h1>
            Our <span> Teams </span>{" "}
          </h1>
        </div>
        <div className="main-content mt-5">
          <Row>
            <Col>
              <Card>
                <div className="profile-detail">
                  <img
                    src={profile}
                    className="avatar img-fluid rounded-circle"
                    alt="logo"
                  />
                  <Row className="pt-2">
                    <Col>
                      <span>Mochammad Andi Rambana</span>
                    </Col>
                  </Row>
                  <Row>
                    <Col>
                      <span>As Frontend Engineer</span>
                    </Col>
                  </Row>
                </div>
              </Card>
            </Col>
            <Col>
              <Card>
                <div className="profile-detail">
                  <img
                    src={profile}
                    className="avatar img-fluid rounded-circle"
                    alt="logo"
                  />
                  <Row className="pt-2">
                    <Col>
                      <span>Muhammad Rifqi Setiawan</span>
                    </Col>
                  </Row>
                  <Row>
                    <Col>
                      <span>As Frontend Engineer</span>
                    </Col>
                  </Row>
                </div>
              </Card>
            </Col>
          </Row>
          {/* 2 */}
          <Row>
            <Col className="pt-5">
              <Card>
                <div className="profile-detail">
                  <img
                    src={profile}
                    className="avatar img-fluid rounded-circle"
                    alt="logo"
                  />
                  <Row className="pt-2">
                    <Col>
                      <span>Navis Abdullah Farhan</span>
                    </Col>
                  </Row>
                  <Row>
                    <Col>
                      <span>As Backend Developer</span>
                    </Col>
                  </Row>
                </div>
              </Card>
            </Col>
            <Col className="pt-5">
              <Card>
                <div className="profile-detail">
                  <img
                    src={profile}
                    className="avatar img-fluid rounded-circle"
                    alt="logo"
                  />
                  <Row className="pt-2">
                    <Col>
                      <span>Paskahl herbert Simarmata</span>
                    </Col>
                  </Row>
                  <Row>
                    <Col>
                      <span>As Backend Developer</span>
                    </Col>
                  </Row>
                </div>
              </Card>
            </Col>
          </Row>
          {/* 3 */}
          <Row>
            <Col className="pt-5">
              <Card>
                <div className="profile-detail">
                  <img
                    src={profile}
                    className="avatar img-fluid rounded-circle"
                    alt="logo"
                  />
                  <Row className="pt-2">
                    <Col>
                      <span>Maulana Dwi Wahyudi</span>
                    </Col>
                  </Row>
                  <Row>
                    <Col>
                      <span>As Backend Developer</span>
                    </Col>
                  </Row>
                </div>
              </Card>
            </Col>
            <Col className="pt-5">
              <Card>
                <div className="profile-detail">
                  <img
                    src={profile}
                    className="avatar img-fluid rounded-circle"
                    alt="logo"
                  />
                  <Row className="pt-2">
                    <Col>
                      <span>Dwi Robbi Prasetyo</span>
                    </Col>
                  </Row>
                  <Row>
                    <Col>
                      <span>As Backend Developer</span>{" "}
                    </Col>
                  </Row>
                </div>
              </Card>
            </Col>
          </Row>
        </div>
      </Container>
      <Footer />
    </section>
  );
}
