import React from 'react';
import profile from '../../../Assets/fotoProfil.jpg';
// import icon from '../../../Assets/Vector.svg';
import { HeartIcon, ChatAltIcon } from '@heroicons/react/outline';
function HeadContentForum() {
  return (
    <>
      <div className="flex items-center justify-between ">
        <div className="flex gap-3 items-center font-poppins">
          <img
            className="w-10 h-10 object-cover border-2 border-orange-500 rounded-full shadow-md"
            src={profile}
            alt=""
          />
          <div className="">
            <p className="text-md font-semibold">Agung</p>
            <p className="text-xs">Dibuat 7 juni 2022</p>
          </div>
        </div>
        <div>
          <div className='flex items-center gap-6'>
            <div className="flex items-center gap-1">
              <HeartIcon className="h-5 w-5" />
              <p className="">10</p>
            </div>
            <div className="flex items-center gap-1">
              <ChatAltIcon className="h-5 w-5" />
              <p className="">5</p>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default HeadContentForum;
