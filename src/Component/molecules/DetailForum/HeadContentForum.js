import React from 'react';
import profile from '../../../Assets/fotoProfil.png';
// import icon from '../../../Assets/Vector.svg';
import LikeUnlike from '../LikeUnlike';
// import { ChatAltIcon } from '@heroicons/react/outline';
import { useForumStore } from '../../store/ProductStore';
function HeadContentForum() {
  const forum = useForumStore((state) => state.forumId);
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
            <p className="text-md font-semibold">{forum?.user?.username}</p>
            <p className="text-xs">Dibuat {(forum.created_at)?.substring(0, 10)}</p>
          </div>
        </div>
        <div>
          <div className="flex items-center gap-6">
            <LikeUnlike type="likeForum" />
            {/* <div className="flex items-center gap-1">
              <ChatAltIcon className="h-5 w-5 text-red-600" />
              <p className="">{props.answer}</p>
            </div> */}
          </div>
        </div>
      </div>
    </>
  );
}

export default HeadContentForum;
