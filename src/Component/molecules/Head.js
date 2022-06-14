import React from 'react';
import { Link } from 'react-router-dom';

function Head(props) {
  return (
    <>
      <div className=" w-full h-full flex justify-center items-start gap-6 flex-col md:items-center md:justify-between md:flex-row ">
        <div className="flex flex-col md:gap-2 ">
          <p className="font-poppins font-semibold text-3xl">
            {props.titleHead}
          </p>
          <p className="text-md font-poppins text-slate-800">
            {props.contentHead}
          </p>
        </div>
        <Link to={props.path}>
          <button className="shadow-md px-6 py-1.5 bg-orange-500 rounded-lg border-neutral-700 border-2 text-white font-bold hover:bg-orange-600 hover:text-white hover:border-black">
            {props.nameButton}
          </button>
        </Link>
      </div>
      <div className=" flex gap-5 items-center w-full justify-end">
        <input
          type="text"
          placeholder="Cari berdasarkan Tag"
          className="shadow-md font-poppins rounded-2xl text-lg px-3 border-2 w-full max-w-md border-slate-800 bg-transparent h-9"
        />
        <button className="px-6 py-1.5 shadow-lg bg-orange-500 rounded-2xl border-neutral-700 border-2 text-white font-bold hover:bg-orange-600 hover:text-white hover:border-black">
          Telusuri
        </button>
      </div>
    </>
  );
}

export default Head;
