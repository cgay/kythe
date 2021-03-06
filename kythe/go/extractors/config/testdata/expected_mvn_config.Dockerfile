FROM openjdk:8 as java
FROM gcr.io/kythe_repo/kythe-javac-extractor-artifacts as javac-extractor-artifacts
FROM maven:latest as maven

FROM debian:jessie
VOLUME /repo
ENV KYTHE_ROOT_DIRECTORY=/repo
VOLUME /out
ENV KYTHE_OUTPUT_DIRECTORY=/out
WORKDIR /repo
COPY --from=java /docker-java-home /docker-java-home
COPY --from=java /etc/java-8-openjdk /etc/java-8-openjdk
ENV JAVA_HOME=/docker-java-home
ENV PATH=$JAVA_HOME/bin:$PATH
COPY --from=javac-extractor-artifacts /opt/kythe/extractors/javac-wrapper.sh /opt/kythe/extractors/javac-wrapper.sh
COPY --from=javac-extractor-artifacts /opt/kythe/extractors/javac_extractor.jar /opt/kythe/extractors/javac_extractor.jar
COPY --from=javac-extractor-artifacts /opt/kythe/extractors/mvn-extract.sh /opt/kythe/extractors/mvn-extract.sh
COPY --from=javac-extractor-artifacts /opt/kythe/extractors/mvn_pom_preprocessor.jar /opt/kythe/extractors/mvn_pom_preprocessor.jar
ENV REAL_JAVAC=$JAVA_HOME/bin/javac
ENV JAVAC_EXTRACTOR_JAR=/opt/kythe/extractors/javac_extractor.jar
ENV JAVAC_WRAPPER=/opt/kythe/extractors/javac-wrapper.sh
ENV MVN_POM_PREPROCESSOR=/opt/kythe/extractors/mvn_pom_preprocessor.jar
COPY --from=maven /usr/share/maven /usr/share/maven
COPY --from=maven /etc/ssl /etc/ssl
ENV MAVEN_HOME=/usr/share/maven
ENV PATH=$MAVEN_HOME/bin:$PATH
ENTRYPOINT ["/opt/kythe/extractors/mvn-extract.sh"]
