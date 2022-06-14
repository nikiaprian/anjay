import React, { useState, useEffect } from 'react';

function ViewTag() {
  let dataTagStatic = [
    {
      id: 1,
      text: 'Python',
    },
    {
      id: 2,
      text: 'Scripts',
    },
    {
      id: 3,
      text: 'Basic',
    },
  ];
  const [tags, setTags] = useState([]);

  useEffect(() => {
    setTags(dataTagStatic);
    //eslint-disable-next-line
  }, []);

  return (
    <>
      <div className="flex flex-row items-center justify-start gap-3">
        {tags.map((tag) => {
          return (
            <div key={tag.id} className="bg-orange-500 font-poppins  opacity-90 text-white px-2 py-1.5 rounded-md shadow-md">
              #{tag.text}
            </div>
          );
        })}
      </div>
    </>
  );
}

export default ViewTag;
