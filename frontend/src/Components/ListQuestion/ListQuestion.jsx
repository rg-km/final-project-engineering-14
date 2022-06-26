import React from "react";

import { Link } from "react-router-dom";

export default function ListQuestion(question, answer) {
  return (
    <section id="list-question">
      <div className="card">
        <div className="card-body">
          <div className="row card-tittle mb-5 mt-3">
            <div className="col-6">
              <h5 className="card-title float-start ml-3">Soal ke - 1 / 15</h5>
            </div>
            <div className="col-6">
              <Link to="/Level">
                <button type="button" className="btn btn-warning kanan mr-3">
                  Next
                </button>
              </Link>
            </div>
          </div>

          <div className="card-text">
            <p>{question}</p>
            <form action="/action_page.php">
              <div className="form-check">
                <input
                  type="radio"
                  className="form-check-input"
                  id="radio1"
                  name="optradio"
                  value="option1"
                />
                <label className="form-check-label" htmlFor="radio1">
                  {answer}
                </label>
              </div>
              <div className="form-check">
                <input
                  type="radio"
                  className="form-check-input"
                  id="radio2"
                  name="optradio"
                  value="option2"
                />
                <label className="form-check-label" htmlFor="radio2">
                  {answer}
                </label>
              </div>
              <div className="form-check">
                <input
                  type="radio"
                  className="form-check-input"
                  id="radio3"
                  name="optradio"
                  value="option3"
                />
                <label className="form-check-label" htmlFor="radio3">
                  {answer}
                </label>
              </div>
              <button type="submit" className="btn btn-success mt-3">
                Check
              </button>
            </form>
          </div>
        </div>
      </div>
    </section>
  );
}
