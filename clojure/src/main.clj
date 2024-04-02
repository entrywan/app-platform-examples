(ns main
  (:gen-class)
  (:use ring.adapter.jetty))

(defn handler [request]
  {:status 200
   :headers {"Content-Type" "text/plain"}
   :body "hello, clojure!"})

(defn -main
  [& args]
  (run-jetty handler {:port 8001}))
