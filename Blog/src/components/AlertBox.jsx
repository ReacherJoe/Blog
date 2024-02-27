import React, { useContext, useState } from "react";
import { Button, Modal } from "antd";
import { RxCross1 } from "react-icons/rx";
import { AiOutlineMail } from "react-icons/ai";
import { FcGoogle } from "react-icons/fc";
import { useNavigate } from "react-router-dom";
import { useGoogleLogin } from "@react-oauth/google";
import axios from "axios";
import { config } from "../config/config";
import { AppContext } from "../contexts/AppContext";

const AlertBox = ({ open, setOpen, text }) => {
  const { setAppData, ...contextData } = useContext(AppContext);
  const [tab, setTab] = useState(1);
  const navigate = useNavigate();

  const signUp = useGoogleLogin({
    onSuccess: async (tokenResponse) => {
      const token = tokenResponse.access_token;

      const response = await axios.get(
        "https://www.googleapis.com/oauth2/v2/userinfo",
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      if (response.status) {
        const data = {
          name: response.data.name,
          email: response.data.email,
          password: "",
          profile: response.data.picture,
        };
        const res = await axios.post(
          `${config.apiBaseUrl}/auth/google-auth`,
          data
        );

        if (res.status) {
          const accessToken = res.data.accessToken;
          if (accessToken) {
            setAppData({ ...contextData, accessToken });
            localStorage.setItem("accessToken", accessToken);
            setOpen(false);
            return navigate("/");
          }
        }
      }
    },
  });

  return (
    <Modal
      visible={open}
      onCancel={() => setOpen(false)}
      footer={null}
      centered
    >
      {tab === 1 && (
        <div className="text-light-text flex flex-col items-center gap-4 p-5">
          <Button
            className="self-end"
            onClick={() => setOpen(false)}
            icon={<RxCross1 />}
          />

          <h1 className="mb-5 text-xl">{text}</h1>

          <Button
            onClick={signUp}
            className="flex items-center gap-4 py-2 px-4 rounded-full min-w-[250px]"
            type="primary"
          >
            <FcGoogle />
            <span>Sign up with Google</span>
          </Button>

          <Button
            onClick={() => navigate("/signUp")}
            className="flex items-center gap-4 py-2 px-4 rounded-full min-w-[250px]"
            type="primary"
          >
            <AiOutlineMail />
            <span>Sign up with email</span>
          </Button>

          <div className="mt-10">
            Already have an account?{" "}
            <span
              onClick={() => setTab(2)}
              className="font-bold text-blue-500 cursor-pointer"
            >
              Sign In
            </span>
          </div>
        </div>
      )}

      {tab === 2 && (
        <div className="text-light-text flex flex-col items-center gap-4 p-5">
          <Button
            className="self-end"
            onClick={() => setOpen(false)}
            icon={<RxCross1 />}
          />

          <h1 className="mb-5 text-xl">Sign In</h1>

          <Button
            className="flex items-center gap-4 py-2 px-4 rounded-full min-w-[250px]"
            type="primary"
          >
            <FcGoogle />
            <span>Sign in with Google</span>
          </Button>

          <Button
            onClick={() => navigate("/signIn")}
            className="flex items-center gap-4 py-2 px-4 rounded-full min-w-[250px]"
            type="primary"
          >
            <AiOutlineMail />
            <span>Sign in with email</span>
          </Button>

          <div onClick={() => setTab(1)} className="mt-10">
            Don't have an account?{" "}
            <span className="font-bold text-blue-500 cursor-pointer">
              Sign Up
            </span>
          </div>
        </div>
      )}
    </Modal>
  );
};

export default AlertBox;
