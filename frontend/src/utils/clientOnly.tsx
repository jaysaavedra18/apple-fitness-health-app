"use client";

import React, { useState, useEffect } from "react";

const ClientOnly = ({ children }: { children: React.ReactNode }) => {
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    setIsClient(true);
  }, []);

  if (!isClient) {
    return null; // Or a loading placeholder if desired: <p>Loading...</p>
  }

  return <>{children}</>;
};

export default ClientOnly;
