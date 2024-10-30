import axios from "axios";
import { useState } from "react";
import { useForm } from "react-hook-form";
import qs from "qs";

export type FormData = {
  username: string;
  firstName: string;
  lastName: string;
  password: string;
};

const inputClass = "border-2 rounded p-2";
const labelClass = "flex text-xl justify-end items-center w-28 mr-3";

const CreateUserForm = () => {
  const { register, handleSubmit } = useForm<FormData>();
  const [loading, setLoading] = useState(false);
  const onSubmit = async ({
    username,
    password,
    firstName,
    lastName,
  }: FormData) => {
    setLoading(true);

    try {
      // TODO: catch response and go to login page if successful
      await axios.post(
        "http://localhost:61942/user",
        qs.stringify({
          username: username,
          password: password,
          firstName: firstName,
          lastName: lastName,
        }),
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        }
      );
    } catch (error) {
      console.error("Error creating user: ", error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="flex flex-col w-full items-center justify-center gap-y-2"
    >
      <div className="flex">
        <label className={labelClass}>First Name:</label>
        <input className={inputClass} type="text" {...register("firstName")} />
      </div>
      <div className="flex">
        <label className={labelClass}>Last Name:</label>
        <input className={inputClass} type="text" {...register("lastName")} />
      </div>
      <div className="flex">
        <label className={labelClass}>Username:</label>
        <input className={inputClass} type="text" {...register("username")} />
      </div>
      <div className="flex">
        <label className={labelClass}>Password:</label>
        <input
          className={inputClass}
          type="password"
          {...register("password")}
        />
      </div>
      <div className="flex flex-col gap-4 justify-center">
        <div className="flex w-full justify-center">
          <button>Create User</button>
        </div>
      </div>
    </form>
  );
};

export default CreateUserForm;
