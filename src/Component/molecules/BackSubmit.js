import React from 'react';

function BackSubmit() {
  return (
    <>
      <div className="flex items-center justify-end gap-3">
        <button className="shadow-lg px-4 py-1.5 bg-white rounded-md border-orange-500 border-2 text-orange-500 font-bold hover:bg-orange-600 hover:text-white hover:border-neutral-700">
          Kembali
        </button>
        <button className="shadow-lg px-6 py-1.5 bg-orange-500 rounded-md border-neutral-700 border-2 text-white font-bold hover:bg-orange-600 hover:text-white hover:border-orange-500">
          Kirim
        </button>
      </div>
    </>
  );
}

export default BackSubmit;
