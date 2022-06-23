import React, { useEffect, useContext } from "react";
import { useNavigate, Outlet } from "react-router-dom";
import AuthContext from "../store/auth-context";
// import { useSelector } from "react-redux";

function PrivateRoute() {
	const user = useContext(AuthContext);

	// console.log("token = ", user);
	const navigate = useNavigate();
	// let location = useLocation();
	useEffect(() => {
		const role = localStorage.getItem("role");

		if (!user.role === "admin" && !role) {
			// console.log("admin login");
			navigate("/HomePage");
		}
	}, []);
	return (
		<div>
			<Outlet />
		</div>
	);
}

export default PrivateRoute;