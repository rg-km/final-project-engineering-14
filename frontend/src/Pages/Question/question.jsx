import React from "react";
import "./question.css";

const question = () => {
	return (
		<section id="question-page">
			<div className="container">
				<h1>Question</h1>
				<div className="card">
					<div className="card-body">
						<div className="row card-tittle mb-5 mt-3">
							<div className="col-6">
								<h5 className="card-title float-start ml-3">
									Soal ke - 1 / 15
								</h5>
							</div>
							<div className="col-6">
								<button type="button" className="btn btn-warning kanan mr-3">
									Next
								</button>
							</div>
						</div>

						<div className="card-text">
							<p>
								Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec
								euismod, nisi vel consectetur euismod, nisi nisl consectetur
								nisi, eget consectetur nisi nisl eget consectetur nisi nisl eget
								consectetur nisi nisl eget consectetur nisi nisl eget
								consectetur nisi nisl eget consectetur nisi nisl eget
								consectetur nisi nisl eget
							</p>
							<form action="/action_page.php">
								<div class="form-check">
									<input
										type="radio"
										class="form-check-input"
										id="radio1"
										name="optradio"
										value="option1"
									/>
									<label class="form-check-label" for="radio1">
										Option 1
									</label>
								</div>
								<div class="form-check">
									<input
										type="radio"
										class="form-check-input"
										id="radio2"
										name="optradio"
										value="option2"
									/>
									<label class="form-check-label" for="radio2">
										Option 2
									</label>
								</div>
								<div class="form-check">
									<input
										type="radio"
										class="form-check-input"
										id="radio3"
										name="optradio"
										value="option3"
									/>
									<label class="form-check-label">Option 3</label>
								</div>
								<button type="submit" class="btn btn-success mt-3">
									Check
								</button>
							</form>
						</div>
					</div>
				</div>
			</div>
		</section>
	);
};

export default question;
