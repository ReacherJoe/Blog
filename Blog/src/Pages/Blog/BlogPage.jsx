import React, { useState } from "react";
import { defaultStyles } from "../styles/styles";
import { Alert } from "antd";
import { BlogCard } from "./BlogCard";
import { BlogCardSkeleton } from "../../skeletons/BlogCardSkeleton";
import { ImageSlider } from "../../components/Slider";
import { useQuery } from "@tanstack/react-query";
import { getInfinitePosts, getPopularPosts } from "../../api/posts";
import Skeleton from "react-loading-skeleton";
import { Post } from "../../typings/types";

export const BlogPage = () => {
  const [page, setPage] = useState(Number(sessionStorage.getItem("page")) || 1);
  const [open, setOpen] = useState(false);
  const popularPostsQuery = useQuery({
    queryKey: ["popularPosts"],
    queryFn: getPopularPosts,
  });

  const { isLoading, isError, error, data } = useQuery({
    queryKey: ["posts", page],
    queryFn: () => getInfinitePosts({ pageParam: page }),
  });

  if (!isLoading && !data?.data) return <div>no post yet!</div>;
  if (isError) return <h1>{JSON.stringify(error)}</h1>;

  return (
    <div
      className={`${defaultStyles.theme} ${defaultStyles.padding}  py-8 grid gap-x-5 gap-y-10   place-items-center  grid-cols-[repeat(auto-fit,minmax(300px,1fr))] 	`}
    >
      {!popularPostsQuery.isLoading ? (
        <ImageSlider popularPosts={popularPostsQuery.data} />
      ) : (
        <div className="w-full col-span-full  aspect-[10/5]">
          <Skeleton className="w-full h-full " />
        </div>
      )}
      {!isLoading ? (
        data?.data.map((post) => {
          return <BlogCard key={post.id} post={post} />;
        })
      ) : (
        <BlogCardSkeleton cards={6} />
      )}

      <div className="flex items-center gap-5 overflow-hidden rounded-md col-span-full">
        <button
          disabled={page === 1}
          className="px-4 py-2 rounded-md border-[1px] border-dark-accent1 text-dark-accent1 disabled:border-gray-400 disabled:text-gray-400"
          onClick={() => {
            setPage((old) => Math.max(old - 1, 0));
            sessionStorage.setItem("page", String(Math.max(page - 1, 0)));
          }}
        >
          PREV
        </button>
        <span>{`Page ${page} of ${data ? data.totalPages : "0"}`}</span>
        <button
          disabled={
            !data?.totalPages || page === data?.totalPages ? true : false
          }
          className="px-4 py-2 rounded-md border-[1px] border-dark-accent1 text-dark-accent1 disabled:border-gray-400 disabled:text-gray-400"
          onClick={() => {
            setPage((old) => old + 1);
            sessionStorage.setItem("page", String(page + 1));
          }}
        >
          NEXT
        </button>
      </div>

      <Alert
        message="Please signIn to like!!"
        type="warning"
        showIcon
        closable
        onClose={() => setOpen(false)}
        className={open ? "" : "hidden"}
      />
    </div>
  );
};
