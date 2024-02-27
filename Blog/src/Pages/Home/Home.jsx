import { googleLogout } from "@react-oauth/google";

export const HomePage = () => {
  return (
    <div className="mx-auto">
      <button onClick={googleLogout}>Sign out</button>
    </div>
  );
};
