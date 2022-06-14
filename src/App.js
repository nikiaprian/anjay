import React from 'react';
import './App.css';
import './output.css';
import Spiner from './Assets/Spinners/Spiner';

//Routing
import { Routes, Route } from 'react-router-dom';
import HomePage from './Component/pages/HomePage';
const DetailForumPage = React.lazy(() =>
  import('./Component/pages/DetailForumPage')
);
const NotFound = React.lazy(() => import('./Component/pages/NotFound'));
const DetailBlogPage = React.lazy(() =>
  import('./Component/pages/DetailBlogPage')
);
const CreateBlogPage = React.lazy(() =>
  import('./Component/pages/CreateBlogPage')
);
const BlogPage = React.lazy(() => import('./Component/pages/BlogPage'));
const ForumPage = React.lazy(() => import('./Component/pages/ForumPage'));
const CreateForumPage = React.lazy(() =>
  import('./Component/pages/CreateForumPage')
);

function App() {
  return (
    <>
      <Routes>
        <Route
          path="/"
          element={
            <React.Suspense
              fallback={
                <div>
                  <Spiner />
                </div>
              }
            >
              <HomePage />
            </React.Suspense>
          }
        />

        {/* forumPage */}
        <Route path="forumin">
          <Route
            index
            element={
              <React.Suspense
                fallback={
                  <div>
                    <Spiner />
                  </div>
                }
              >
                <ForumPage />
              </React.Suspense>
            }
          />
          <Route
            path="createforum"
            element={
              <React.Suspense
                fallback={
                  <div>
                    <Spiner />
                  </div>
                }
              >
                <CreateForumPage />
              </React.Suspense>
            }
          />
          <Route
            path="detailforum/:idforum"
            element={
              <React.Suspense
                fallback={
                  <div>
                    <Spiner />
                  </div>
                }
              >
                <DetailForumPage />
              </React.Suspense>
            }
          />
        </Route>

        {/* blogPage */}
        <Route path="blogin">
          <Route
            index
            element={
              <React.Suspense
                fallback={
                  <div>
                    <Spiner />
                  </div>
                }
              >
                <BlogPage />
              </React.Suspense>
            }
          />
          <Route
            path="createblog"
            element={
              <React.Suspense
                fallback={
                  <div>
                    <Spiner />
                  </div>
                }
              >
                <CreateBlogPage />
              </React.Suspense>
            }
          />
          <Route
            path="detailblog/:idblog"
            element={
              <React.Suspense
                fallback={
                  <div>
                    <Spiner />
                  </div>
                }
              >
                <DetailBlogPage />
              </React.Suspense>
            }
          />
        </Route>

        <Route
          path="*"
          element={
            <React.Suspense
              fallback={
                <div>
                  <Spiner />
                </div>
              }
            >
              <NotFound />
            </React.Suspense>
          }
        />
      </Routes>
    </>
  );
}

export default App;
