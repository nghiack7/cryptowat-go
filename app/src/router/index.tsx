import type { RouteObject } from "react-router-dom";
import Layout from "../components/Layout";
import HomePage from "../pages/home.page";
import LoginPage from "../pages/login.page";
import ProfilePage from "../pages/profile.page";
import RegisterPage from "../pages/register.page";
import EthPage from "../pages/eth.page";
import EthPosition from "../pages/eth.position.page";

// @ts-ignore
const normalRoutes: RouteObject = {
  path: "*",
  element: <Layout />,
  children: [
    {
      index: true,
      element: <HomePage />,
    },
    {
      path: "profile",
      element: <ProfilePage />,
    },
    {
      path: "login",
      element: <LoginPage />,
    },
    {
      path: "register",
      element: <RegisterPage />,
    },
    {
      path: "eth",
      element: <EthPage />,
    },
    {
      path: "eth/positions",
      element: <EthPosition />
    },
  ],
};

const routes: RouteObject[] = [normalRoutes];

export default routes;
