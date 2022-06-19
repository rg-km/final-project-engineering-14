import React from "react";
import "./style.css";
import Structure from "../../../Assets/Images/LearningPath-Struct.png";

export default function LearningPath() {
  return (
    <section id="learning-path">
      <div className="path">
        <h2>Learning Path Backend</h2>
        <div className="laerning-path">
          <img src={Structure} className="img-fluid" alt="learning-path" />
        </div>
      </div>
    </section>
  );
}
