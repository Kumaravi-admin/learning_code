import { useContext } from "react";
import { UserContext } from "./Hook";

export default function Message() {
  const user = useContext(UserContext);

  return <h1>Hello user : {user}</h1>;
}
