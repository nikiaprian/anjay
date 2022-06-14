import React from 'react';
import ViewTag from '../ViewTag';
import {Link} from 'react-router-dom';
function CardForum(props) {
  return (
    <>
      <div className="border-2 px-5 bg-[#f2f2f2] border-gray-300 w-full rounded-md shadow-lg flex flex-col gap-3">
      <Link to={`/forumin/detailforum/${props.id}`}>
        <div className="flex mt-2 ml-2 gap-2  pr-8 md:pr-0">
          <div className="flex flex-col pt-10 pr-5">
            <div className="flex flex-col items-center ">
              <p className="font-bold text-xl text-orange-500">{props.like}</p>
              <p className="font-medium text-lg">Suka</p>
            </div>
            <div className="flex flex-col items-center  p-2">
              <p className="font-bold text-xl text-orange-500">
                {props.answer}
              </p>
              <p className="font-medium text-lg">Jawaban</p>
            </div>
          </div>
          <div className="w-10/12 xl:w-full pt-10  flex flex-col gap-2">
            <p className="font-bold text-lg font-poppins">{props.title}</p>
            <p className="text-sm font-poppins truncate w-11/12">
              {props.content}
            </p>
            <div className="">
              <ViewTag />
            </div>
          </div>
        </div>
        <div className="flex mb-5 mr-5 justify-end gap-3 items-center font-poppins">
          <img
            className="w-10 h-10 object-cover border-2 border-orange-500 rounded-full shadow-md"
            src={props.profileImg}
            alt=""
          />
          <div className="flex flex-col">
            <p className="text-md font-semibold">{props.user}</p>
            <p className="text-xs">{props.date}</p>
          </div>
        </div>
        </Link>
      </div>
    </>
  );
}

export default CardForum;
